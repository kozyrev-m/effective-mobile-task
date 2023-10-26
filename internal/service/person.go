package service

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
	"github.com/kozyrev-m/effective-mobile-task/internal/store"
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
func (svc *Service) FindPersonByID(ctx context.Context, id uint64) (*entities.Person, error) {
	person, err := svc.store.GetPersonByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// DeletePerson deletes person by id.
func (svc *Service) DeletePerson(ctx context.Context, personID uint64) (uint64, error) {
	id, err := svc.store.DeletePerson(ctx, personID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdatePerson updates person.
func (svc *Service) UpdatePerson(ctx context.Context, personID uint64, params entities.Person) (*entities.Person, error) {
	_, err := svc.store.GetPersonByID(ctx, personID)
	if err != nil {
		return nil, err
	}

	return svc.store.UpdatePerson(ctx, personID, params)
}

// GetPersons gets persons by filter.
func (svc *Service) GetPersons(ctx context.Context, filter store.Filter) ([]*entities.Person, bool, error) {
	var hasNextPage bool = false

	persons, err := svc.store.GetPersons(ctx, filter)
	if err != nil {
		return nil, false, err
	}

	if len(persons) > *filter.PerPage {
		persons = persons[:len(persons)-1]
		hasNextPage = true
	}

	return persons, hasNextPage, nil
}
