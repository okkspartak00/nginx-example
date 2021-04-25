package repository

import (
	"kitchen-service/model"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	Database *gorm.DB
}

func (repo *RestaurantRepository) CreateRestaurant(restaurant *model.Restaurant) error {
	result := repo.Database.Create(restaurant)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}
