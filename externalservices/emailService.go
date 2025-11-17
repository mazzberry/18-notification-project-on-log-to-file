package externalservices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
}

func (e *EmailService) SendEmail(order *entities.Order) {
	fmt.Printf("Email sent : %v\n", order)
}

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email sent to receiver: %s \n Message: %s  \n", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
