package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
	"go.uber.org/mock/gomock"
)

var (
	errTestCreatePerson = errors.New("Can't create person")
	errTestFindPerson   = errors.New("Can't find person")
	errTestDeletePerson = errors.New("Can't delete person")
	errTestUpdatePerson = errors.New("Can't update person")
)

func (suite *serviceTestSuite) TestServiceCreatePerson() {
	testCases := []struct {
		name        string
		expectedErr error
		initMock    func() entities.Person
	}{
		{
			name:        "success",
			expectedErr: nil,
			initMock: func() entities.Person {
				// init person for test
				person := entities.TestPerson()

				// init mock for current test case
				suite.mock.EXPECT().CreatePerson(gomock.Any(), person).Return(uint64(1), nil)

				return person
			},
		},
		{
			name:        "fail",
			expectedErr: errTestCreatePerson,
			initMock: func() entities.Person {
				// init person for test
				person := entities.TestPerson()

				// init mock for current test case
				suite.mock.EXPECT().CreatePerson(gomock.Any(), person).Return(uint64(0), errTestCreatePerson)

				return person
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			person := tc.initMock()
			_, err := suite.svc.CreatePerson(context.TODO(), person)
			suite.Equal(tc.expectedErr, err)
		})
	}
}

func (suite *serviceTestSuite) TestServiceFindPerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name        string
		expectedErr error
		initMock    func() uint64
	}{
		{
			name:        "success",
			expectedErr: nil,
			initMock: func() uint64 {
				// init person id for test
				id := rnd.Intn(32)

				// init person data for test
				person := entities.TestPerson()
				person.ID = uint64(id)

				// init mock for current test case
				suite.mock.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(&person, nil)

				return uint64(id)
			},
		},
		{
			name:        "fail",
			expectedErr: errTestFindPerson,
			initMock: func() uint64 {
				// init person id for test
				personID := rnd.Intn(32)

				// init mock for current test case
				suite.mock.EXPECT().GetPersonByID(gomock.Any(), uint64(personID)).Return(nil, errTestFindPerson)

				return uint64(personID)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			personID := tc.initMock()
			_, err := suite.svc.FindPersonByID(context.TODO(), personID)
			suite.Equal(tc.expectedErr, err)
		})
	}
}

func (suite *serviceTestSuite) TestServiceDeletePerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name        string
		expectedErr error
		initMock    func() uint64
	}{
		{
			name:        "success",
			expectedErr: nil,
			initMock: func() uint64 {
				// init person id for test
				personID := rnd.Intn(32)

				// init mock for current test case
				suite.mock.EXPECT().DeletePerson(gomock.Any(), uint64(personID)).Return(uint64(personID), nil)

				return uint64(personID)
			},
		},
		{
			name:        "faile",
			expectedErr: errTestDeletePerson,
			initMock: func() uint64 {
				// init person id for test
				personID := rnd.Intn(32)

				// init mock for current test case
				suite.mock.EXPECT().DeletePerson(gomock.Any(), uint64(personID)).Return(uint64(0), errTestDeletePerson)

				return uint64(personID)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			id := tc.initMock()
			_, err := suite.svc.DeletePerson(context.TODO(), id)
			suite.Equal(tc.expectedErr, err)
		})
	}
}

func (suite *serviceTestSuite) TestServiceUpdatePerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name        string
		expectedErr error
		initMock    func() (uint64, entities.Person)
	}{
		{
			name:        "success",
			expectedErr: nil,
			initMock: func() (uint64, entities.Person) {
				// init person id for test
				id := rnd.Intn(32)

				// init person data for test
				person := entities.TestPerson()
				person.ID = uint64(id)

				// init params for updating
				newName := "Sergey"
				newAge := 45
				newNationality := "USA"
				params := entities.Person{Name: &newName, Age: &newAge, Nationality: &newNationality}

				// init updated person data for mock
				updatedPerson := entities.TestPerson()
				updatedPerson.ID = uint64(id)
				updatedPerson.Name = &newName
				updatedPerson.Age = &newAge
				updatedPerson.Nationality = &newNationality

				// init mock for current test case
				suite.mock.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(&person, nil)
				suite.mock.EXPECT().UpdatePerson(gomock.Any(), uint64(id), params).Return(&updatedPerson, nil)

				return uint64(id), params
			},
		},
		{
			name:        "fail",
			expectedErr: errTestUpdatePerson,
			initMock: func() (uint64, entities.Person) {
				// init person id for test
				id := rnd.Intn(32)

				// init person data for test
				person := entities.TestPerson()
				person.ID = uint64(id)

				// init params for updating
				newName := "Sergey"
				newAge := 45
				newNationality := "USA"
				params := entities.Person{Name: &newName, Age: &newAge, Nationality: &newNationality}

				// init mock for current test case
				suite.mock.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(&person, nil)

				suite.mock.EXPECT().UpdatePerson(gomock.Any(), uint64(id), params).Return(nil, errTestUpdatePerson)

				return uint64(id), params
			},
		},
		{
			name:        "no person",
			expectedErr: errTestFindPerson,
			initMock: func() (uint64, entities.Person) {
				// init person id for test
				id := rnd.Intn(32)

				// init mock for current test case
				suite.mock.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(nil, errTestFindPerson)

				// init data for updating
				newName := "Sergey"
				newAge := 45
				newNationality := "USA"
				params := entities.Person{Name: &newName, Age: &newAge, Nationality: &newNationality}

				return uint64(id), params
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			id, params := tc.initMock()
			_, err := suite.svc.UpdatePerson(context.TODO(), id, params)
			suite.Equal(tc.expectedErr, err)
		})
	}
}
