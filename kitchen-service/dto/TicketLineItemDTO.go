package dto

type TicketLineItemDTO struct {
	MenuItemId   string `json:"menuItemId"`
	MenuItemName string `json:"menuItemName"`
	Quantity     int    `json:"quantity"`
}
