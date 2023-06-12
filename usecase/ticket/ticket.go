package ticket

import (
	"log"

	googleD "github.com/mohfahrur/interop-service-a/domain/google"
	interopcD "github.com/mohfahrur/interop-service-a/domain/interopc"
	"github.com/mohfahrur/interop-service-a/entity"
)

type TicketAgent interface {
	SendEmail(user string, item string) (err error)
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

func (uc *TicketUsecase) SendEmail(entity.SendEmailRequest) (err error) {
	// uc.googleDomain.CreateMessage()
	// uc.googleDomain.SendEmail()
	log.Println("send status to interopC")
	return
}
