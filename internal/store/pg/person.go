package pg

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

// sqlCreatePerson is used to create person.
//
//	$1 - name
//	$2 - patronymic
//	$3 - surname
//	$4 - age
//	$5 - gender
//	$6 - nationality
//
// It return person id.
var sqlCreatePerson = `
	INSERT INTO persons (name, patronymic, surname, age, gender, nationality)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
`

// CreatePerson creates person.
func (s *Store) CreatePerson(ctx context.Context, p entities.Person) (uint64, error) {
	// id - id of created person
	var id uint64

	err := s.conn.QueryRowContext(
		ctx, sqlCreatePerson,
		p.Name, p.Patronymic, p.Surname, p.Age, p.Gender, p.Nationality,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
