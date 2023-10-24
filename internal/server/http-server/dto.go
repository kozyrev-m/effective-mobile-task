package httpserver

import "github.com/kozyrev-m/effective-mobile-task/internal/entities"

// PersonCreateRequest uses as a struct to store data from request body to create person.
type PersonBodyRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// convert converts person dto to entity.
func convert(p PersonBodyRequest) entities.Person {
	return entities.Person{
		Name:        &p.Name,
		Patronymic:  &p.Patronymic,
		Surname:     &p.Surname,
		Age:         &p.Age,
		Gender:      &p.Gender,
		Nationality: &p.Nationality,
	}
}
