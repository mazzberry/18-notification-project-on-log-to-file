package externalservices

import "fmt"

type NilNotifyService struct {
}

func (e *NilNotifyService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("nilNotifyService: Receiver: %s, Message: %s", receiver, message)
	return nil
}

func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
