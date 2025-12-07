package externalservices

import (
	"fmt"
	"notification/core"
	"notification/entities"

	"github.com/rs/zerolog"
)

var logger = core.NewFileLogger()

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

	fmt.Printf("nilNotifyService: Receiver: %s, Message: %s", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("Message info", zerolog.Dict().
			Str("receiver", receiver).Str("receiver", receiver)).
		Msg("Message sent:")

	return nil
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
