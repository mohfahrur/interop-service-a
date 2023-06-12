package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	googleD "github.com/mohfahrur/interop-service-a/domain/google"
	interopcD "github.com/mohfahrur/interop-service-a/domain/interopc"
	"github.com/mohfahrur/interop-service-a/entity"
	ticketUC "github.com/mohfahrur/interop-service-a/usecase/ticket"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	credentialsFile, err := ioutil.ReadFile(pwd + "/credential.json")
	if err != nil {
		log.Println(err)
		return
	}

	googleDomain := googleD.NewGoogleDomain(credentialsFile)
	interopcDomain := interopcD.NewinteropcDomain()
	ticketUsecase := ticketUC.NewTicketUsecase(*googleDomain, *interopcDomain)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from service a",
		})
	})
	r.POST("/send-email", func(c *gin.Context) {
		var sendEmailReq entity.SendEmailRequest
		err = c.BindJSON(&sendEmailReq)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		err = ticketUsecase.SendEmail(sendEmailReq)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.Run(":5000")
}
