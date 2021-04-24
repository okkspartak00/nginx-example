package service

import (
	"github.com/veljkomaksimovic/nginx-example/model"
	"github.com/veljkomaksimovic/nginx-example/repository"
)

type ConsumerService struct {
	Repo *repository.ConsumerRepository
}

func (service *ConsumerService) CreateConsumer(consumer *model.Consumer) error {
	service.Repo.CreateConsumer(consumer)
	return nil
}
