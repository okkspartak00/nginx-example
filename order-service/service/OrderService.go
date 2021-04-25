package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/veljkomaksimovic/order-service/model"
	"github.com/veljkomaksimovic/order-service/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (service *OrderService) CreateOrder(order *model.Order) error {
	url := fmt.Sprintf("http://%s:%s/verify/%s", os.Getenv("CONSUMER_SERVICE_DOMAIN"), os.Getenv("CONSUMER_SERVICE_PORT"), order.ConsumerID)
	print(url)
	resp, err := http.Get(url)
	if err != nil {
		print(err)
		return err
	}
	if resp.StatusCode == 404 {
		return fmt.Errorf("consumer with id %s does not exist", order.ConsumerID)
	}
	//TODO check if kitchen exists
	order.OrderStatus = model.PENDING
	service.Repo.CreateOrder(order)
	return nil
}

func (service *OrderService) ChangeStatus(orderId string, status string) error {
	id, err := uuid.Parse(orderId)
	if err != nil {
		print(err)
		return err
	}
	var orderStatus int
	switch status {
	case "pending":
		orderStatus = model.PENDING
	case "accepted":
		orderStatus = model.ACCEPTED
	case "rejected":
		orderStatus = model.REJECTED
	}
	service.Repo.UpdateOrder(id, orderStatus)
	return nil
}
