package service

import (
	"chiapitest/model"
	"chiapitest/repo"
	"errors"
)

func GetPeople() []model.Person {
	return repo.GetAllPeople()
}

func GetPerson(personID string) model.Person {
	return repo.FindPersonByID(personID)
}

func AddPerson(person model.PersonRequest) error {
	newPerson := model.Person{
		ID:        "123",
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Address:   person.Address,
	}
	err := repo.CreatePerson(newPerson)
	return errors.New("Unable to create person error: '" + err.Error() + "'.")
}

func RemovePerson(personID string) error {
	person := repo.FindPersonByID(personID)
	if (person == model.Person{}) {
		return errors.New("Unable to find person with ID '" + personID + "'")
	}

	return repo.DeletePerson(person)
}
