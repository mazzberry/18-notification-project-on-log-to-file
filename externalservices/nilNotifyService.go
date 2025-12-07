package externalservices

import (
	"fmt"

	"github.com/rs/zerolog"
)

type NilNotifyService struct {
}

func (e *NilNotifyService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("nilNotifyService: Receiver: %s, Message: %s", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("Message info", zerolog.Dict().
			Str("receiver", receiver).Str("receiver", receiver)).
		Msg("Message sent:")
	return nil
}

func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
