package service

import (
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/errs"
)

type IEmailBlastService interface {
	SendEmail(request domain.EmailBlast) *errs.Exception
}

type DefaultEmailBlastService struct {
	repo domain.IEmailBlastRepository
}

func NewEmailBlastService(repository domain.IEmailBlastRepository) DefaultEmailBlastService {
	return DefaultEmailBlastService{repository}
}

func (s DefaultEmailBlastService) SendEmail(request domain.EmailBlast) *errs.Exception {
	err := s.repo.Send(request)

	if err != nil {
		return err
	}

	return nil
}
