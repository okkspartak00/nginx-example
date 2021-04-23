package controller

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/veljkomaksimovic/nginx-example/model"
	"github.com/veljkomaksimovic/nginx-example/poststore"
)

var PS *poststore.PostStore

func Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var consumer model.Consumer
	err := json.NewDecoder(r.Body).Decode(&consumer)
	if err != nil {
		fmt.Println("test")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(consumer)
	ret_consumer, err := PS.Post(&consumer)
	if err != nil {
		fmt.Println("test2")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*ret_consumer)
}

func Verify() {
	//TODO
}
