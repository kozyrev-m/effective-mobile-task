package store

// Filter defines filter to select persons.
type Filter struct {
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
