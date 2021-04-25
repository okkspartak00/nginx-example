package handler

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/veljkomaksimovic/nginx-example/model"
	"github.com/veljkomaksimovic/nginx-example/service"
)

type ConsumerHandler struct {
	Service *service.ConsumerService
}

func (handler *ConsumerHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

func (handler *ConsumerHandler) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	var consumer model.Consumer
	err := json.NewDecoder(r.Body).Decode(&consumer)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(consumer)
	err = handler.Service.CreateConsumer(&consumer)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func Verify() {
	//TODO
}
