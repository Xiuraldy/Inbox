package usecase

import (
	repository "challenge/repository/email"
	"fmt"
)

type UseCase interface {
	GetEmail(from int, search string) ([]interface{}, error)
}

type useCaseEmail struct {
	RepositoryEmail repository.Repository
}

func NewUseCaseEmail(repositoryEmail repository.Repository) UseCase {
	return &useCaseEmail{
		RepositoryEmail: repositoryEmail,
	}
}

func (u *useCaseEmail) GetEmail(from int, search string) ([]interface{}, error) {
	hits, err := u.RepositoryEmail.GetEmail(from, search)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails from repository: %w", err)
	}

	return hits, nil
}
