package service

import (
	"errors"
	"test-rod/repository"
)

//go:generate mockery --name=Service --output=../mock/mock_service --outpkg=mock_service
type Service interface {
	ValidUser(string) error
}

type ServiceSt struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *ServiceSt {
	return &ServiceSt{
		repository: repository,
	}
}

func (u ServiceSt) ValidUser(user string) error {

	if user == "" {
		return errors.New("invalid user")
	}

	if _, err := u.repository.GetUser(user); err != nil {
		return err
	}

	return nil
}
