// Package integration is used to receive additional data about a person and to add this data to a person.
package integration

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

// Integration is interfeice to get additional person data and to set it to person.
type Integration interface {
	ReceiveAndSet(ctx context.Context, person entities.Person) (*entities.Person, error)
}
