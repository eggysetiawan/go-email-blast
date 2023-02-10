package domain

import (
	"github.com/eggysetiawan/go-email-blast/errs"
	"github.com/go-mail/mail"
	"os"
	"strconv"
)

type DefaultRepositorySmtp struct {
}

func (d DefaultRepositorySmtp) Send(request EmailBlast) *errs.Exception {
	m := mail.NewMessage()

	m.SetHeader("From", "rahmat.setiawan@lawencon.com")

	m.SetHeader("To", request.Email)

	m.SetHeader("Subject", request.Subject)

	m.SetBody("text/html", request.Body)

	if request.Filename != "" {
		m.Attach(request.Filename)
	}

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	dialer := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	if err := dialer.DialAndSend(m); err != nil {

		return errs.NewUnexpectedException(err.Error())

	}

	return nil
}

func NewEmailBlastRepositorySmtp() DefaultRepositorySmtp {
	return DefaultRepositorySmtp{}
}
