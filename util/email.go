package util

import (
	"gopkg.in/gomail.v2"
)

type SmtpInfo struct {
	Host     string
	Port     int
	User     string
	Passwrod string
	From     string
}

func Send(to, subject, body string, smtp SmtpInfo) error {
	//fmt.Printf("%s,%s,%s", to, subject, body)
	m := gomail.NewMessage()
	m.SetHeader("From", smtp.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtp.Host, smtp.Port, smtp.User, smtp.Passwrod)

	err := d.DialAndSend(m)
	return err
}
