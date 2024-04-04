package common

import (
	"medioker-bank/config"
	"medioker-bank/model/dto"

	"gopkg.in/gomail.v2"
)

type Mailer interface {
	SendEmail(payload dto.Mail) error
}

type mailer struct {
	cfg config.MailerConfig
}

func (m *mailer) SendEmail(payload dto.Mail) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.cfg.MailerUsername)
	message.SetHeader("To", payload.Receiver)
	message.SetHeader("Subject", payload.Subject)
	message.SetBody("text/html", payload.Body)

	d := gomail.NewDialer(m.cfg.MailerHost, m.cfg.MailerPort, m.cfg.MailerUsername, m.cfg.MailerPassword)

	if err := d.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

func NewMailer(cfg config.MailerConfig) Mailer {
	return &mailer{cfg: cfg}
}
