package usecase

import (
	repository "challenge/repository/total"
	"fmt"
)

type UseCase interface {
	GetTotal(from int, search string) ([]interface{}, error)
}

type useCaseTotal struct {
	RepositoryTotal repository.Repository
}

func NewUseCaseTotal(repositoryTotal repository.Repository) UseCase {
	return &useCaseTotal{
		RepositoryTotal: repositoryTotal,
	}
}

func (u *useCaseTotal) GetTotal(from int, search string) ([]interface{}, error) {
	hits, err := u.RepositoryTotal.GetTotal(from, search)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails from repository: %w", err)
	}

	return hits, nil
}
