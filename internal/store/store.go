// Package store works with storage app.
package store

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

// RepositoryPerson describes person repository.
type RepositoryPerson interface {
	// GetPersonByID returns person by id
	GetPersonByID(ctx context.Context, personID uint64) (*entities.Person, error)
	// DeletePerson deletes person by id
	DeletePerson(ctx context.Context, personID uint64) (uint64, error)
	// UpdatePerson edites person
	UpdatePerson(ctx context.Context, personID uint64, person entities.Person) (*entities.Person, error)
	// CreatePerson creates person
	CreatePerson(ctx context.Context, person entities.Person) (uint64, error)
}

// Store describes an abstract storage.
type Store interface {
	RepositoryPerson
}
