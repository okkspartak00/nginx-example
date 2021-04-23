package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/veljkomaksimovic/nginx-example/controller"
	ps "github.com/veljkomaksimovic/nginx-example/poststore"
)

func handleFunc() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.Hello).Methods("GET")
	router.HandleFunc("/", controller.CreateBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	store, err := ps.New()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	controller.PS = store
	handleFunc()
}
