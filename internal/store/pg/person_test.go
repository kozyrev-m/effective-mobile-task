package pg

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

func (suite *storeTestSuite) TestCreatePerson() {
	suite.NotNil(suite.store)

	person1 := entities.Person{Name: "Ivan", Patronymic: "Ivanovich", Surname: "Ivanov", Age: 30, Gender: "man", Nationality: "Russian"}
	id, err := suite.store.CreatePerson(context.TODO(), person1)
	suite.NoError(err)
	suite.NotZero(id)

	person2 := entities.Person{Name: "Sergey", Patronymic: "Sergeyvich", Surname: "Sergeev", Age: 29, Gender: "man", Nationality: "Russian"}
	id, err = suite.store.CreatePerson(context.TODO(), person2)
	suite.NoError(err)
	suite.NotZero(id)
}

func (suite *storeTestSuite) TestGetPersonByID() {
	suite.NotNil(suite.store)

	_, err := suite.store.GetPersonByID(context.TODO(), 1234)
	suite.Error(err)

	person := entities.Person{Name: "Ivan", Patronymic: "Ivanovich", Surname: "Ivanov", Age: 30, Gender: "man", Nationality: "Russian"}
	id, err := suite.store.CreatePerson(context.TODO(), person)
	suite.NoError(err)
	person.ID = id

	found, err := suite.store.GetPersonByID(context.TODO(), id)
	suite.NoError(err)
	suite.NotNil(found)
	suite.Equal(person, *found)
}

func (suite *storeTestSuite) TestDeletePerson() {
	suite.NotNil(suite.store)

	_, err := suite.store.DeletePerson(context.TODO(), 1)
	suite.Error(err)

	person := entities.Person{Name: "Ivan", Patronymic: "Ivanovich", Surname: "Ivanov", Age: 30, Gender: "man", Nationality: "Russian"}
	id, err := suite.store.CreatePerson(context.TODO(), person)
	suite.NoError(err)
	suite.NotZero(id)

	deletedID, err := suite.store.DeletePerson(context.TODO(), id)
	suite.NoError(err)
	suite.Equal(id, deletedID)
}
