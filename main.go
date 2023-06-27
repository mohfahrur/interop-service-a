package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	googleD "github.com/mohfahrur/interop-service-a/domain/google"
	interopcD "github.com/mohfahrur/interop-service-a/domain/interopc"
	"github.com/mohfahrur/interop-service-a/entity"
	ticketUC "github.com/mohfahrur/interop-service-a/usecase/ticket"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	log.SetFlags(log.Llongfile)
	email := os.Getenv("email")
	password := os.Getenv("password")

	googleDomain := googleD.NewGoogleDomain(email, password)
	interopcDomain := interopcD.NewinteropcDomain()
	ticketUsecase := ticketUC.NewTicketUsecase(*googleDomain, *interopcDomain)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from service a",
		})
		return
	})
	r.POST("/send-email", func(c *gin.Context) {
		var sendEmailReq entity.SendEmailRequest
		err := c.BindJSON(&sendEmailReq)
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
		return
	})
	r.Run(":5000")
}
