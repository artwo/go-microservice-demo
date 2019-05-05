package service

import (
	"chiapitest/models"
	"chiapitest/repos"
	"errors"
)

func GetPeople() []models.Person {
	return repos.GetAllPeople()
}

func GetPerson(personID string) models.Person {
	return repos.FindPersonByID(personID)
}

func AddPerson(person models.PersonRequest) error {
	newPerson := models.Person{
		ID:        "123",
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Address:   person.Address,
	}
	err := repos.CreatePerson(newPerson)
	return errors.New("Unable to create person error: '" + err.Error() + "'.")
}

func RemovePerson(personID string) error {
	person := repos.FindPersonByID(personID)
	if (person == models.Person{}) {
		return errors.New("Unable to find person with ID '" + personID + "'")
	}

	return repos.DeletePerson(person)
}
