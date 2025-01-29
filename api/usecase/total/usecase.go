package usecase

import (
	repository "challenge/repository/total"
	"errors"
	"fmt"
)

// Define la interfaz para los casos de uso relacionados con el total de emails
type UseCase interface {
	GetTotal(from int, search string) ([]interface{}, error) // Obtiene el número total de correos que coincidan con los parámetros
}

// Implementación de la lógica de negocio para el total de emails
type useCaseTotal struct {
	RepositoryTotal repository.Repository
}

// Crea una nueva instancia del caso de uso para total
func NewUseCaseTotal(repositoryTotal repository.Repository) UseCase {
	return &useCaseTotal{
		RepositoryTotal: repositoryTotal,
	}
}

// Obtiene el número total de correos según los parámetros de entrada
func (u *useCaseTotal) GetTotal(from int, search string) ([]interface{}, error) {
	// Validar el parámetro "from"
	if from < 0 {
		return nil, errors.New("invalid 'from' parameter: must be >= 0")
	}

	// Delegar la consulta al repositorio
	hits, err := u.RepositoryTotal.GetTotal(from, search)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails from repository: %w", err)
	}

	// Retornar mensaje si no se encontraron resultados
	if len(hits) == 0 {
		return nil, errors.New("no total emails found for the given query")
	}

	return hits, nil
}
