package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID           uuid.UUID        `json:"id"`
	TicketState  TicketState      `json:"ticketState"`
	Items        []TicketLineItem `json:"items"`
	RestaurantID uuid.UUID        `json:"restaurant"`
}

func (ticket *Ticket) BeforeCreate(store *gorm.DB) error {
	ticket.ID = uuid.New()
	return nil
}
