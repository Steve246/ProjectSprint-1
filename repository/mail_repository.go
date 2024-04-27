package repository

import (
	"7Zero4/config"
	"log"

	"gopkg.in/gomail.v2"
)

type MailRepository interface {
	SendMail(receiverMail string, body string) error
}

type mailRepository struct {
	mailConfig config.MailConfig
}

func (m *mailRepository) SendMail(receiverMail string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", m.mailConfig.CONFIG_SENDER_NAME)
	mailer.SetHeader("To", receiverMail)
	// mailer.SetAddressHeader("Cc", "alamt email pengirim", "isi teks) // cc
	mailer.SetHeader("Subject", "OTP Spendy")
	mailer.SetBody("text/html", body)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		m.mailConfig.CONFIG_SMTP_HOST,
		m.mailConfig.CONFIG_SMTP_PORT,
		m.mailConfig.CONFIG_AUTH_EMAIL,
		m.mailConfig.CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Mail sent!")
	return nil
}

func NewMailRepository(mailConfig config.MailConfig) MailRepository {
	repo := new(mailRepository)
	repo.mailConfig = mailConfig
	return repo
}
