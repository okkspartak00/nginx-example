package handler

import (
	"fmt"
	"kitchen-service/service"
	"net"
	"net/http"
)

type KitchenHandler struct {
	Service *service.KitchenService
}

func (handler *KitchenHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}
