// Package svc contains service layer that works with business logic.
package svc

import "github.com/kozyrev-m/effective-mobile-task/internal/store"

// Service is where all business logic should happen.
type Service struct {
	store store.Store
}

// NewService creates service.
func NewService(s store.Store) *Service {
	return &Service{
		store: s,
	}
}
