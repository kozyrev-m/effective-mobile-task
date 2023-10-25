package agent

// ageResponseBody - use to get age.
type ageResponseBody struct {
	Count int    `json:"count"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
}

// genderResponseBody - use to get gender.
type genderResponseBody struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

// nationalizeResponseBody - use to nationalize.
type nationalizeResponseBody struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []country `json:"country"`
}

// country - it contains country id and probability.
type country struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
