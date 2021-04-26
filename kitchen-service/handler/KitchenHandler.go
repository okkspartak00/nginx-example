package handler

import (
	"encoding/json"
	"fmt"
	"kitchen-service/dto"
	"kitchen-service/service"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type KitchenHandler struct {
	Service *service.TicketService
}

func (handler *KitchenHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

func (handler *KitchenHandler) Verify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifying...")
	vars := mux.Vars(r)
	var items dto.TicketLineItemsDTO
	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(items)
	result := handler.Service.Verify(vars["restaurantId"], items)
	w.Header().Set("Content-Type", "application/json")
	if !result {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *KitchenHandler) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var items dto.TicketLineItemsDTO
	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(items)
	if handler.Service.Create(vars["restaurantId"], vars["orderId"], items) {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (handler *KitchenHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ticketId"]
	status := vars["state"]
	if !handler.Service.TicketRepo.ExistsById(id) {
		w.WriteHeader(http.StatusBadRequest)
	}
	err := handler.Service.Update(id, status)
	if err != nil {
		print(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
