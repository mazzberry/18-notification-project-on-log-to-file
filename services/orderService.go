package services

import (
	"fmt"
	"notification/core"
	"notification/entities"
	"notification/externalservices"
)

var logger = core.NewFileLogger()

type OrderService struct {
	Notifier externalservices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Order created : %v", order)
	logger.Info().Interface("order", order).Msgf("Order created.")
	orderService.Notifier = externalservices.NewNotifier(order.NotificationType)

	orderService.Notifier.SendNotify(order.UserId, "Order created")

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
