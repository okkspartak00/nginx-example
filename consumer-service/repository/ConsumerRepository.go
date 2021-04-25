package repository

import (
	"fmt"

	"github.com/veljkomaksimovic/nginx-example/model"
	"gorm.io/gorm"
)

type ConsumerRepository struct {
	Database *gorm.DB
}

func (repo *ConsumerRepository) CreateConsumer(consumer *model.Consumer) error {
	result := repo.Database.Create(consumer)
	//TODO convert to logs
	print(result.Error.Error())
	fmt.Println(result.RowsAffected)
	return nil
}
