package services

import (
	"fmt"
	"notification/entities"
	"notification/externalservices"
)

type OrderService struct {
	Notifier externalservices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Order created : %v", order)
	orderService.Notifier = externalservices.NewNotifier(order.NotificationType)

	orderService.Notifier.SendNotify(order.UserId, "Order created")
	

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
