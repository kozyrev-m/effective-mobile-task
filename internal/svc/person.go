package svc

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

// CreatePerson is method that creates person.
func (svc *Service) CreatePerson(ctx context.Context, person entities.Person) (uint64, error) {
	id, err := svc.store.CreatePerson(ctx, person)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// FindPersonByID is method that find person by id.
func (c *Service) FindPersonByID(ctx context.Context, id uint64) (*entities.Person, error) {
	person, err := c.store.GetPersonByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// DeletePerson deletes person by id.
func (c *Service) DeletePerson(ctx context.Context, personID uint64) (uint64, error) {
	id, err := c.store.DeletePerson(ctx, personID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdatePerson updates person.
func (c *Service) UpdatePerson(ctx context.Context, personID uint64, params entities.Person) (*entities.Person, error) {
	_, err := c.store.GetPersonByID(ctx, personID)
	if err != nil {
		return nil, err
	}

	return c.store.UpdatePerson(ctx, personID, params)
}
