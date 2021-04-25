package service

import (
	"fmt"
	"kitchen-service/dto"
	"kitchen-service/model"
	"kitchen-service/repository"

	"github.com/google/uuid"
)

type TicketService struct {
	MenuItemRepo   *repository.MenuItemRepository
	RestaurantRepo *repository.RestaurantRepository
	TicketRepo     *repository.TicketRepository
}

func (service *TicketService) Verify(restaurantID string, items dto.TicketLineItemsDTO) bool {
	if !service.RestaurantRepo.ExistsById(restaurantID) {
		fmt.Println("Restaurant does not exist!")
		return false
	}
	for _, item := range items.TicketLineItems {
		if !service.MenuItemRepo.ExistsByIdAndRestaurantID(item.MenuItemId, restaurantID) {
			fmt.Println("Menu item does not exist in the restaurant!")
			return false
		}
	}
	return true
}

func (service *TicketService) Create(restaurantId string, orderId string, items dto.TicketLineItemsDTO) {
	if !service.RestaurantRepo.ExistsById(restaurantId) {
		fmt.Println(("Restaurant does not exist!"))
	} else {
		fmt.Println(("Restaurant found"))
		orderUuid, _ := uuid.Parse(orderId)

		restaurantUuid, _ := uuid.Parse(restaurantId)
		var it []model.TicketLineItem
		ticket := model.Ticket{ID: orderUuid, RestaurantID: restaurantUuid, TicketState: model.PENDING, Items: it}
		for _, item := range items.TicketLineItems {
			menuItem := service.MenuItemRepo.FindById(item.MenuItemId)
			fmt.Println("1")
			ticketLineItem := model.TicketLineItem{MenuItem: menuItem, Quantity: item.Quantity}
			ticket.AddItem(ticketLineItem)

		}
		fmt.Println(ticket)
		service.TicketRepo.CreateTicket(&ticket)

	}
}

func (service *TicketService) Update(restaurantId string, orderId string, items dto.TicketLineItemsDTO) {
	//TODO
}
