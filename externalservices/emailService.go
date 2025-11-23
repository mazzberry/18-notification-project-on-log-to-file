package externalservices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
}

func (e *EmailService) SendEmail(order *entities.Order) error {
	fmt.Printf("Email sent : %v\n", order)
	return nil
}

func (e *EmailService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("Email sent to receiver: %s \n Message: %s  \n", receiver, message)
	return nil
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
