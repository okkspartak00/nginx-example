package main

import (
	"kitchen-service/model"
	"kitchen-service/repository"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initRestaurantRepo(database *gorm.DB) *repository.RestaurantRepository {
	return &repository.RestaurantRepository{Database: database}
}

func initDB() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("restaurant.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Restaurant{})

	/*Loading test data*/
	restaurant := model.Restaurant{Name: "Trattoria"}

	database.Create(&restaurant)

	tickets := []model.Ticket{
		{TicketState: 0, RestaurantID: restaurant.ID},
		{TicketState: 1, RestaurantID: restaurant.ID},
	}

	for _, ticket := range tickets {
		database.Create(&ticket)
	}

	return database
}

func main() {
	database := initDB()
	initRestaurantRepo(database)
}
