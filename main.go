// main --> orderService -> emailService || smsService
package main

import (
	"notification/entities"
	//"notification/externalservices"
	"notification/services"
)

func main() {
	order1 := entities.Order{
		Id:               1,
		UserFullName:     "John Doe",
		UserId:           "09134212531",
		Price:            2000,
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
	orderService.CreateOrder(&order1)

	orderService.CreateOrder(&order2)

	orderService.CreateOrder(&order3)

}
