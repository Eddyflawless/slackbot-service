package sms

import (
	"fmt"
)

type MessageService interface {
	SendChargeNotification(int) error
}

type SMSService struct{}

type MyService struct {
	messageService MessageService
}

func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Send production charge notification")
	return nil
}

func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil

}

func main() {
	fmt.Println("sms main-->")

	smsService := SMSService{}

	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
