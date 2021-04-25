package dto

import "github.com/google/uuid"

type TicketLineItemDTO struct {
	MenuItemId   uuid.UUID `json:"menuItemId"`
	MenuItemName string    `json:"menuItemName"`
	Quantity     int       `json:"quantity"`
}
