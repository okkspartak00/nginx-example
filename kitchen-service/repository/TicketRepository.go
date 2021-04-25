package repository

import (
	"fmt"
	"kitchen-service/model"

	"gorm.io/gorm"
)

type TicketRepository struct {
	Database *gorm.DB
}

func (repo *TicketRepository) ExistsById(ticketID string) bool {

	if err := repo.Database.First(&model.Ticket{}, "id = ?", ticketID).Error; err != nil {
		return false
	}
	return true
}

func (repo *TicketRepository) FindById(ticketID string) *model.Ticket {

	ticket := &model.Ticket{}
	repo.Database.First(&ticket, "id = ?", ticketID)
	return ticket
}

func (repo *TicketRepository) CreateTicket(ticket *model.Ticket) error {
	result := repo.Database.Create(ticket)
	//TODO convert to logs
	print(result.Error.Error())
	fmt.Println(result.RowsAffected)
	fmt.Println("Ticket Created")
	return nil
}
