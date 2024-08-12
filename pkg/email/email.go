package email

import (
	"errors"
	"gopkg.in/gomail.v2"
)

var _smtpMail *SmtpMail

type Address struct {
	Name    string
	Address string
}

type SmtpMail struct {
	Host   string
	Port   int
	User   string
	Pass   string
	From   *Address
	Dialer *gomail.Dialer
}

func (smtp *SmtpMail) Send(to *Address, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(smtp.From.Address, smtp.From.Name))
	m.SetHeader("To", m.FormatAddress(to.Address, to.Name))
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return smtp.Dialer.DialAndSend(m)
}

func OnInit(host, user, pass, address, name string, port int) *SmtpMail {
	_smtpMail = &SmtpMail{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		From: &Address{
			Name:    name,
			Address: address,
		},
	}
	_smtpMail.Dialer = gomail.NewDialer(_smtpMail.Host, _smtpMail.Port, _smtpMail.User, _smtpMail.Pass)
	return _smtpMail
}
func SendMail(to *Address, subject, body string) error {
	if _smtpMail == nil {
		return errors.New("nil")
	}
	return _smtpMail.Send(to, subject, body)
}
