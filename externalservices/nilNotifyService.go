package externalservices

import "fmt"

type NilNotifyService struct {
}

func (e *NilNotifyService) SendNotify(receiver string, message string) {
	fmt.Printf("nilNotifyService: Receiver: %s, Message: %s", receiver, message)
}

func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
