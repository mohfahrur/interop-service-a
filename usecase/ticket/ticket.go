package ticket

import (
	"log"

	googleD "github.com/mohfahrur/interop-service-a/domain/google"
	interopcD "github.com/mohfahrur/interop-service-a/domain/interopc"
)

type TicketAgent interface {
	Buy(user string, item string) (err error)
}

type TicketUsecase struct {
	googleDomain   googleD.GoogleDomain
	interopcDomain interopcD.InteropcDomain
}

func NewTicketUsecase(
	googleD googleD.GoogleDomain,
	interopcD interopcD.InteropcDomain) *TicketUsecase {

	return &TicketUsecase{
		googleDomain:   googleD,
		interopcDomain: interopcD}
}

func (uc *TicketUsecase) UpdateEmailNotificationStatus(user string, status string) (err error) {
	log.Println("send status to interopC")
	return
}
