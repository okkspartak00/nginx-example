package main

import (
	"fmt"
	"kitchen-service/handler"
	"kitchen-service/model"
	"kitchen-service/repository"
	"kitchen-service/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initRestaurantRepo(database *gorm.DB) *repository.RestaurantRepository {
	return &repository.RestaurantRepository{Database: database}
}

func initMenuItemRepository(database *gorm.DB) *repository.MenuItemRepository {
	return &repository.MenuItemRepository{Database: database}
}

func initTicketRepository(database *gorm.DB) *repository.TicketRepository {
	return &repository.TicketRepository{Database: database}
}

func initServices(menuItemRepo *repository.MenuItemRepository, restaurantRepo *repository.RestaurantRepository, ticketRepo *repository.TicketRepository) *service.TicketService {
	return &service.TicketService{MenuItemRepo: menuItemRepo, RestaurantRepo: restaurantRepo, TicketRepo: ticketRepo}
}

func initHandler(service *service.TicketService) *handler.KitchenHandler {
	return &handler.KitchenHandler{Service: service}
}

func initDB() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("kitchen.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Restaurant{})
	database.AutoMigrate(&model.MenuItem{})
	database.AutoMigrate(&model.Ticket{})
	database.AutoMigrate(&model.TicketLineItem{})

	/*Loading test data*/

	return database
}

func handleFunc(handler *handler.KitchenHandler) {
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("server running ")

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/create/{restaurantId}/{orderId}", handler.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	database := initDB()
	restaurantRepo := initRestaurantRepo(database)
	menuItemRepo := initMenuItemRepository(database)
	TicketRepository := initTicketRepository(database)
	service := initServices(menuItemRepo, restaurantRepo, TicketRepository)
	handler := initHandler(service)

	// ticketLineItems := []model.TicketLineItem{
	// 	{MenuItem: menuItems[0], Quantity: 2},
	// 	{MenuItem: menuItems[1], Quantity: 2},
	// }

	//restaurant := model.Restaurant{Name: "Trattoria", MenuItems: menuItems}
	restaurant := model.Restaurant{Name: "Trattoria"}
	database.Create(&restaurant)

	menuItems := []model.MenuItem{
		{Name: "Pizza", Restaurant: restaurant},
		{Name: "Pasta", Restaurant: restaurant},
	}

	for _, menuItem := range menuItems {
		database.Create(&menuItem)
	}
	// fmt.Println(menuItemRepo.ExistsByIdAndRestaurantID(menuItems[0].ID, restaurant.ID))

	handleFunc(handler)
}
