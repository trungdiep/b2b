package tasks

import (
	"example/web-service-gin/models"
	"fmt"
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
)

var paymentModels = new(models.PaymentModel)

type SendMail struct {
}

func mailRemainder(email_to []string) email.Email {
	return email.Email{
		To:      email_to,
		From:    "Jordan Wright <test@gmail.com>",
		Subject: "Remind",
		Text:    []byte("You have not paid the full amount"),
		HTML:    []byte("<h1>You have not paid the full amount!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
}

func (e SendMail) Run() {
	merchants, err := paymentModels.GetEmails()
	var emails []string

	if err == nil {
		for _, e := range merchants {
			emails = append(emails, e.Email)
		}
	}

	email := mailRemainder(emails)
	err = email.Send(
		os.Getenv("MAIL_ADDR"),
		smtp.PlainAuth("", os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"), os.Getenv("MAIL_HOST")),
	)

	fmt.Printf("send_mail%v", err)

}
