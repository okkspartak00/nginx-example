package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketLineItem struct {
	ID         uuid.UUID `json:"id"`
	MenuItemID uuid.UUID `json:"item"`
	Quantity   string    `json:"quantity" gorm:"not null"`
	TicketId   string    `json:"ticket"`
}

func (ticketLineItem *TicketLineItem) BeforeCreate(scope *gorm.DB) error {
	ticketLineItem.ID = uuid.New()
	return nil
}
