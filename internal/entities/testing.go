package entities

var (
	name        string = "Ivan"
	patronymic  string = "Ivanovich"
	surname     string = "Ivanov"
	age         int    = 30
	gender      string = "male"
	nationality string = "RU"
)

func TestPerson() Person {
	return Person{
		Name: &name, Patronymic: &patronymic, Surname: &surname,
		Age: &age, Gender: &gender, Nationality: &nationality,
	}
}
