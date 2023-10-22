// Package store works with storage app.
package store

import "context"

// RepositoryPerson describes person repository.
type RepositoryPerson interface {
	// GetPersonByID returns person by id
	GetPersonByID(ctx context.Context, personID uint64) (*Person, error)
	// DeletePerson deletes person by id
	DeletePerson(ctx context.Context, personID uint64) (uint64, error)
	// UpdatePerson edites person
	UpdatePerson(ctx context.Context, person Person) (*Person, error)
	// CreatePerson creates person
	CreatePerson(ctx context.Context, person Person) error
}

// Store describes an abstract storage.
type Store interface {
	RepositoryPerson
}

// Person describes man.
type Person struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}
