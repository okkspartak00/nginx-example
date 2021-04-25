package repository

import (
	"kitchen-service/model"

	"gorm.io/gorm"
)

type MenuItemRepository struct {
	Database *gorm.DB
}

func (repo *MenuItemRepository) CreateMenuItem(menuItem *model.MenuItem) error {
	result := repo.Database.Create(menuItem)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}
