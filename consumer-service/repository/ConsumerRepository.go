package repository

import (
	"github.com/veljkomaksimovic/nginx-example/model"
	"gorm.io/gorm"
)

type ConsumerRepository struct {
	Database *gorm.DB
}

func (repo *ConsumerRepository) CreateConsumer(consumer *model.Consumer) error {
	result := repo.Database.Create(consumer)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}
