package service

import (
	"20241107/class/1/model"
	"20241107/class/1/repository"
)

type OrderService struct {
	OrderRepo repository.Order
}

func InitOrderService(repo repository.Order) *OrderService {
	return &OrderService{OrderRepo: repo}
}

func (orderService OrderService) Create(order *model.Order) error {
	return orderService.OrderRepo.Create(order)
}
