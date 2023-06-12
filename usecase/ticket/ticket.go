package ticket

import (
	"fmt"
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

func (uc *TicketUsecase) SendEmail(req entity.SendEmailRequest) (err error) {
	status := "success"
	msg := fmt.Sprintf("terima kasih %s sudah melakukan pembelian tiket film %s",
		req.User,
		req.Item)

	err = uc.googleDomain.SendEmail(req.Email,
		"notifikasi pembelian",
		msg)
	if err != nil {
		log.Println(err)
		status = "fail"
	}
	uc.interopcDomain.UpdateEmailNotificationStatus(req.User, status)
	return
}
