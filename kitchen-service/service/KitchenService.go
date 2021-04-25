package service

import (
	"kitchen-service/model"
	"kitchen-service/repository"
)

type KitchenService struct {
	MenuItemRepo   *repository.MenuItemRepository
	RestaurantRepo *repository.RestaurantRepository
}

func (service *KitchenService) CreateRestaurant(restaurant *model.Restaurant) error {
	service.RestaurantRepo.CreateRestaurant(restaurant)
	return nil
}

func (service *KitchenService) CreateMenuItem(menuItem *model.MenuItem) error {
	service.MenuItemRepo.CreateMenuItem(menuItem)
	return nil
}
