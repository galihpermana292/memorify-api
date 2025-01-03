package email

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

type Gomail struct {
	dialer *gomail.Dialer
}

func NewMailClient(host, email, password string, port int) *Gomail {

	fmt.Println("email", host, port, email, password)

	return &Gomail{
		gomail.NewDialer(host, port, email, password),
	}
}

func (g *Gomail) SendEmailPaymentNotification(sender, subject, paymentID, username string) error {
	recipients := []string{"basalamah76@gmail.com"}

	message := gomail.NewMessage()
	message.SetHeader("From", sender)
	message.SetHeader("To", recipients...)
	message.SetHeader("Subject", subject)

	var body bytes.Buffer
	path := "pkg/provider/email/sendnotification.html" // TODO: change path
	t, err := template.ParseFiles(path)
	if err != nil {
		return errors.New("error parse html template")
	}

	t.Execute(&body, struct {
		PaymentID string
		Username  string
	}{
		PaymentID: paymentID,
		Username:  username,
	})
	message.SetBody("text/html", body.String())

	if err := g.dialer.DialAndSend(message); err != nil {
		return err
	}
	fmt.Println("success mail sent")
	return nil
}
