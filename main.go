// main --> orderService -> emailService || smsService
package main

import (
	"notification/core"
	"notification/entities"

	//"notification/externalservices"
	"notification/services"
)

var logger = core.NewFileLogger()

func main() {
	order1 := entities.Order{
		Id:               1,
		UserFullName:     "John Doe",
		UserId:           "09134212531",
		Price:            120,
		Status:           true,
		NotificationType: entities.Email,
	}

	order2 := entities.Order{
		Id:               2,
		UserFullName:     "mahbod akbari",
		UserId:           "09132341231",
		Price:            4000,
		Status:           true,
		NotificationType: entities.Sms,
	}

	order3 := entities.Order{
		Id:               2,
		UserFullName:     "masoud rajavi",
		UserId:           "099990341231",
		Price:            4000,
		Status:           true,
		NotificationType: entities.Nil,
	}

	orderService := services.NewOrderService()
	err, _ := orderService.CreateOrder(&order1)
	if err != nil {
		logger.Error().Interface("entity info", order1).Err(err).Msg("order is not valid")
	}
	err, _ = orderService.CreateOrder(&order2)
	if err != nil {
		logger.Error().Interface("entity info", order2).Err(err).Msg("order is not valid")
	}
	err, _ = orderService.CreateOrder(&order3)
	if err != nil {
		logger.Error().Interface("entity info", order3).Err(err).Msg("order is not valid")
	}
}
