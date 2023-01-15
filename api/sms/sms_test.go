package sms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
	mock.Mock
}

// mocked service method
func (m *smsServiceMock) SendChargeNotification(value int) error {

	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)

	args := m.Called(value)

	// it returns whatever we tell it

	return args.Error(0)
}

// stub out every method defined in that interface

func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// create SmsService Mock
func TestChargeCustomer(t *testing.T) {

	smsService := new(smsServiceMock)

	// define what should be retured from SendCharfeNotification
	// return true when 100 is passed to it.
	smsService.On("SendChargeNotification", 100).Return(true)

	myService := MyService{smsService}

	myService.ChargeCustomer(100)
}
