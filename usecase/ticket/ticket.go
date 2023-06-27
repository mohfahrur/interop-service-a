package ticket

import (
	"fmt"
	"log"
	"time"

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
	currentTime := time.Now().Format("Monday, 02 January 2006 15:04")
	msg := fmt.Sprintf("Terima kasih %s telah melakukan pembelian tiket film %s\ntanggal pembelian: %s",
		req.User,
		req.Item,
		currentTime)

	err = uc.googleDomain.SendEmail(req.Email,
		"Notifikasi Pembelian",
		msg)
	if err != nil {
		log.Println(err)
		status = "fail"
	}
	uc.interopcDomain.UpdateEmailNotificationStatus(req.User, status)
	return
}
