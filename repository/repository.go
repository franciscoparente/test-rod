package repository

import "errors"

//go:generate mockery --name=Repository --output=../mock/mock_repository --outpkg=mock_repository
type Repository interface {
	GetUser(string) (string, error)
}

type RepostiroySt struct {
}

func (u RepostiroySt) GetUser(user string) (string, error) {

	if user == "fparente" {
		return "root", nil
	}

	return "", errors.New("not found")
}
