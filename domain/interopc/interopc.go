package interopc

import (
	"log"
)

type InteropcAgent interface {
	UpdateEmailNotificationStatus(user string, status string) error
}

type InteropcDomain struct {
}

func NewinteropcDomain() *InteropcDomain {

	return &InteropcDomain{}
}

func (d *InteropcDomain) UpdateEmailNotificationStatus(user string, status string) (err error) {
	log.Println("send status to interopC")
	return
}
