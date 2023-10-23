package entities

var (
	name        string = "Ivan"
	patronymic  string = "Ivanovich"
	surname     string = "Ivanov"
	age         int    = 30
	gender      string = "man"
	nationality string = "Russian"
)

func TestPerson() Person {
	return Person{
		Name: &name, Patronymic: &patronymic, Surname: &surname,
		Age: &age, Gender: &gender, Nationality: &nationality,
	}
}
