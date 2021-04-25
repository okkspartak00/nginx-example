package main

import (
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

func initServices(menuItemRepo *repository.MenuItemRepository, restaurantRepo *repository.RestaurantRepository) *service.KitchenService {
	return &service.KitchenService{MenuItemRepo: menuItemRepo, RestaurantRepo: restaurantRepo}
}

func initHandler(service *service.KitchenService) *handler.KitchenHandler {
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

	menuItems := []model.MenuItem{
		{Name: "Pizza"},
		{Name: "Pasta"},
	}

	restaurant := model.Restaurant{Name: "Trattoria", MenuItems: menuItems}

	ticketLineItems := []model.TicketLineItem{
		{MenuItem: menuItems[0], Quantity: 2},
		{MenuItem: menuItems[1], Quantity: 2},
	}

	database.Create(&restaurant)

	for _, ticketLineItem := range ticketLineItems {
		database.Create(&ticketLineItem)
	}

	return database
}

func handleFunc(handler *handler.KitchenHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	database := initDB()
	restaurantRepo := initRestaurantRepo(database)
	menuItemRepo := initMenuItemRepository(database)
	service := initServices(menuItemRepo, restaurantRepo)
	handler := initHandler(service)
	handleFunc(handler)
}
