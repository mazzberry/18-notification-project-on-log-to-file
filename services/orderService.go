package services

import (
	"errors"
	"fmt"
	"notification/core"
	"notification/entities"
	"notification/externalservices"
)

var logger = core.NewFileLogger()

type OrderService struct {
	Notifier externalservices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) (error, *entities.Order) {

	if order.Status == false {
		return errors.New("Order is not valid"), order
	}

	if order.Price < 1000 {
		return errors.New("Order Price is not valid"), order
	}

	fmt.Printf("Order created : %v", order)

	logger.Info().Interface("order", order).Msgf("Order created.")
	
	orderService.Notifier = externalservices.NewNotifier(order.NotificationType)

	logger.Info().Msgf("Notifier founded: %T", orderService.Notifier)

	err := orderService.Notifier.SendNotify(order.UserId, "Order created")

	return err ,order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
