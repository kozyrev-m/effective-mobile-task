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
