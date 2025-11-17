package externalservices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
}

func (e *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("\nSms sent : %v\n", order)
}

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("\nSms sent to receiver: %s \n Message: %s  \n", receiver, message)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
