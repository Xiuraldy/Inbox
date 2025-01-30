package usecase

import (
	repository "challenge/repository/email"
	"errors"
	"fmt"
)

// Define la interfaz para los casos de uso relacionados con emails
type UseCase interface {
	GetEmail(from int, search string) ([]interface{}, error) // Obtiene una lista de correos que coincidan con los parámetros
}

// Implementación de la lógica de negocio para emails
type useCaseEmail struct {
	RepositoryEmail repository.Repository
}

// Crea una nueva instancia del caso de uso para emails
func NewUseCaseEmail(repositoryEmail repository.Repository) UseCase {
	if repositoryEmail == nil {
		panic("repositoryEmail must not be nil")
	}
	return &useCaseEmail{
		RepositoryEmail: repositoryEmail,
	}
}

// Obtiene una lista de correos según los parámetros de entrada
func (u *useCaseEmail) GetEmail(from int, search string) ([]interface{}, error) {
	// Validar el parámetro "from"
	if from < 0 {
		return nil, errors.New("invalid 'from' parameter: must be >= 0")
	}

	// Delegar la consulta al repositorio
	hits, err := u.RepositoryEmail.GetEmail(from, search)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails from repository: %w", err)
	}

	return hits, nil
}
