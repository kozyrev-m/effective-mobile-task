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

// sqlFindPersonByID is used to get person.
//
//	$1 - id
//
// It returns id, name, patronymic, surname, age, gender, nationality.
var sqlFindPersonByID = `
	SELECT id, name, patronymic, surname, age, gender, nationality
	FROM persons
	WHERE id = $1
`

// GetPersonByID finds person by id.
func (s *Store) GetPersonByID(ctx context.Context, personID uint64) (*entities.Person, error) {
	person := &entities.Person{}
	err := s.conn.QueryRowContext(ctx, sqlFindPersonByID, personID).
		Scan(
			&person.ID, &person.Name, &person.Patronymic,
			&person.Surname, &person.Age, &person.Gender,
			&person.Nationality,
		)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// sqlDeletePerson is used to delete row with person by id.
//
//	$1 - id
//
// It returns deleted person id.
var sqlDeletePersonByID = `
	DELETE FROM persons WHERE id = $1 RETURNING id
`

// DeletePerson deletes person.
func (s *Store) DeletePerson(ctx context.Context, personID uint64) (uint64, error) {
	var deletedPersonID uint64
	err := s.conn.QueryRowContext(ctx, sqlDeletePersonByID, personID).Scan(&deletedPersonID)
	if err != nil {
		return 0, err
	}

	return deletedPersonID, nil
}

// sqlUpdatePerson is used to update person.
//
//	$1 - id
//
//	$2 - name
//	$3 - patronymic
//	$4 - surname
//	$5 - age
//	$6 - gender
//	$7 - nationality
var sqlUpdatePerson = `
	UPDATE persons
	SET
		name = COALESCE($2, name),
		patronymic = COALESCE($3, patronymic),
		surname = COALESCE($4, surname),
		age = COALESCE($5, age),
		gender = COALESCE($6, gender),
		nationality = COALESCE($7, nationality),
		updated_at = now()
	WHERE id = $1
	RETURNING id, name, patronymic, surname, age, gender, nationality
`

// UpdatePerson updates person.
func (s *Store) UpdatePerson(ctx context.Context, personID uint64, p entities.Person) (*entities.Person, error) {
	person := &entities.Person{}
	err := s.conn.QueryRowContext(
		ctx,
		sqlUpdatePerson,
		personID,
		p.Name, p.Patronymic, p.Surname,
		p.Age, p.Gender, p.Nationality,
	).Scan(
		&person.ID, &person.Name, &person.Patronymic,
		&person.Surname, &person.Age, &person.Gender,
		&person.Nationality,
	)

	if err != nil {
		return nil, err
	}

	return person, nil
}
