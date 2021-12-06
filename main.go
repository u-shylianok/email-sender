package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.POST("/send", Sender)

	router.Run(":8080")
}

type Message struct {
	Reciever string `json:"reciever"`
	Text     string `json:"text"`
}

func Sender(c *gin.Context) {
	var message Message
	if err := c.BindJSON(&message); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	result, err := SendEmail(message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func SendEmail(message Message) (string, error) {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// Receiver email address.
	// to := []string{
	// 	"email@example.com",
	// }
	to := []string{
		message.Reciever,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("html/template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		TemplateVar1 string
		TemplateVar2 string
	}{
		TemplateVar1: "Some template message in var 1",
		TemplateVar2: "Sended message from api: " + message.Text,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Email Sent!")

	// Just return result message for example
	return body.String(), err
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Print(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
