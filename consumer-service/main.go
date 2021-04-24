package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/veljkomaksimovic/nginx-example/handler"
	"github.com/veljkomaksimovic/nginx-example/model"
	"github.com/veljkomaksimovic/nginx-example/repository"
	"github.com/veljkomaksimovic/nginx-example/service"
)

func initDB() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("consumers.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Consumer{})

	/*Loading test data*/
	consumers := []model.Consumer{
		{Email: "petar.petrovic@mail.cc", Password: "petar", Name: "petar", Surname: "petrovic"},
		{Email: "ivan.ivanovic@example.cc", Password: "ivan", Name: "ivan", Surname: "ivanovic"},
	}
	for _, consumer := range consumers {
		database.Create(&consumer)
	}
	return database
}

func initRepo(database *gorm.DB) *repository.ConsumerRepository {
	return &repository.ConsumerRepository{Database: database}
}

func initServices(repo *repository.ConsumerRepository) *service.ConsumerService {
	return &service.ConsumerService{Repo: repo}
}

func initHandler(service *service.ConsumerService) *handler.ConsumerHandler {
	return &handler.ConsumerHandler{Service: service}
}
func handleFunc(handler *handler.ConsumerHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreateConsumer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}
