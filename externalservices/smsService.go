package externalservices

import (
	"fmt"
	"notification/entities"

	"github.com/rs/zerolog"
)

type SmsService struct {
}

func (e *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("\nSms sent : %v\n", order)
}

func (e *SmsService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("\nSms sent to receiver: %s \n Message: %s  \n", receiver, message)

	fmt.Printf("nilNotifyService: Receiver: %s, Message: %s", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("Message info", zerolog.Dict().
			Str("receiver", receiver).Str("receiver", receiver)).
		Msg("Message sent:")

	return nil
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
