package repos

import (
	"chiapitest/models"
	"chiapitest/utils"
	"errors"
	"log"
)

var people []models.Person


func init() {
	people = append(
		people,
		models.Person{
			ID: "2",
			Firstname: "Koko",
			Lastname: "Doe",
			Address: &models.Address{
				City: "City Z", State: "State Y"}},
		models.Person{
			ID: "1",
			Firstname: "John",
			Lastname: "Doe",
			Address: &models.Address{
				City: "City X", State: "State X"}})

	log.Printf(utils.ToString(people))
}

func GetAllPeople() [] models.Person {
	return people
}

func FindPersonByID(ID string) models.Person {
	for _, item := range people {
		if item.ID == ID {
			return item
		}
	}
	return models.Person{}
}

func CreatePerson(person models.Person) error {
	people = append(people, person)
	return nil
}

func DeletePerson(person models.Person) error {
	for index, item := range people {
		if item.ID == person.ID {
			people = append(people[:index], people[index + 1:]...)
			return nil
		}
	}
	return errors.New("person does not exist")
}