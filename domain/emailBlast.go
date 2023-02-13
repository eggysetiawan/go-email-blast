package domain

import "github.com/eggysetiawan/go-email-blast/errs"

type EmailBlast struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Attachment string `json:"attachment"`
	Filename   string `json:"filename"`
}

type IEmailBlastRepository interface {
	Send(p EmailBlast) *errs.Exception
}
