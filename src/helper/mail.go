package helper

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func SendEmail() {
	sender := "no-reply@crocodic.net"
	to := []string{
		"laililmahvut@gmail.com",
	}
	user := "b755af4d-1d0d-4c92-ab8a-1b153c6df02f"
	password := "b755af4d-1d0d-4c92-ab8a-1b153c6df02f"

	subject := "Simple HTML mail"
	body := `<p>An old <b>falcon</b> in the sky.</p>`
	request := Mail{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	addr := "smtp.postmarkapp.com:587"
	host := "smtp.postmarkapp.com"

	msg := BuildMessage(request)
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")

}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
