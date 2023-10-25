// Package dto...
package dto

import (
	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
	"github.com/kozyrev-m/effective-mobile-task/internal/store"
)

// PersonsRequestBody
type PersonsRequestBody struct {
	Page        *int    `json:"page"`
	PerPage     *int    `json:"per_page"`
	ID          *uint64 `json:"id"`
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Age         *int    `json:"age"`
	Gender      *string `json:"gender"`
	Nationality *string `json:"nationality"`
}

// Convert converts to filter.
func (p PersonsRequestBody) Convert() store.Filter {
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PerPage == nil {
		defaultPerPage := 10
		p.PerPage = &defaultPerPage
	}

	return store.Filter{
		Page:        p.Page,
		PerPage:     p.PerPage,
		ID:          p.ID,
		Name:        p.Name,
		Surname:     p.Surname,
		Patronymic:  p.Patronymic,
		Age:         p.Age,
		Gender:      p.Gender,
		Nationality: p.Nationality,
	}
}

// PersonsResponseBody.
type PersonsResponseBody struct {
	Persons []*entities.Person `json:"persons"`
}
