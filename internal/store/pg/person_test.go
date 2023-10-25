package pg

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

func (suite *storeTestSuite) TestCreatePerson() {
	suite.NotNil(suite.store)

	person1 := entities.TestPerson()
	id, err := suite.store.CreatePerson(context.TODO(), person1)
	suite.NoError(err)
	suite.NotZero(id)

	person2 := entities.TestPerson()
	*person2.Name = "Sergey"
	*person2.Patronymic = "Sergeyvich"
	*person2.Surname = "Sergeev"
	*person2.Age = 29

	id, err = suite.store.CreatePerson(context.TODO(), person2)
	suite.NoError(err)
	suite.NotZero(id)
}

func (suite *storeTestSuite) TestGetPersonByID() {
	suite.NotNil(suite.store)

	_, err := suite.store.GetPersonByID(context.TODO(), 1234)
	suite.Error(err)

	person := entities.TestPerson()
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

	person := entities.TestPerson()
	id, err := suite.store.CreatePerson(context.TODO(), person)
	suite.NoError(err)
	suite.NotZero(id)

	deletedID, err := suite.store.DeletePerson(context.TODO(), id)
	suite.NoError(err)
	suite.Equal(id, deletedID)
}

func (suite *storeTestSuite) TestUpdatePerson() {
	suite.NotNil(suite.store)

	person := entities.TestPerson()
	id, err := suite.store.CreatePerson(context.TODO(), person)
	suite.NoError(err)

	person.ID = id

	// 1-st case - change some fields
	newName1 := "Mikhail"
	newAge1 := 25
	newParams1 := entities.Person{Name: &newName1, Age: &newAge1}

	_, err = suite.store.UpdatePerson(context.TODO(), id, newParams1)
	suite.NoError(err)

	updated1, _ := suite.store.GetPersonByID(context.TODO(), id)

	// checks of 1-st case
	suite.Equal(newName1, *updated1.Name)
	suite.Equal(newAge1, *updated1.Age)

	// 2-nd case - change all fields
	newName2 := "Ann"
	newPatronymic2 := ""
	newSurname2 := "Smith"
	newAge2 := 23
	newGender2 := "female"
	newNationality2 := "USA"
	newParams2 := entities.Person{
		Name: &newName2, Patronymic: &newPatronymic2, Surname: &newSurname2,
		Age: &newAge2, Gender: &newGender2, Nationality: &newNationality2,
	}

	_, err = suite.store.UpdatePerson(context.TODO(), id, newParams2)
	suite.NoError(err)

	updated2, _ := suite.store.GetPersonByID(context.TODO(), id)

	// checks of 2-nd case
	suite.Equal(newName2, *updated2.Name)
	suite.Equal(newPatronymic2, *updated2.Patronymic)
	suite.Equal(newSurname2, *updated2.Surname)
	suite.Equal(newAge2, *updated2.Age)
	suite.Equal(newGender2, *updated2.Gender)
	suite.Equal(newNationality2, *updated2.Nationality)
}
