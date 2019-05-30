package repo

import (
	"chiapitest/model"
	"chiapitest/utils"
	"errors"
	"log"
)

var people []model.Person

func init() {
	people = append(
		people,
		model.Person{
			ID:        "2",
			Firstname: "Koko",
			Lastname:  "Doe",
			Address: &model.Address{
				City: "City Z", State: "State Y"}},
		model.Person{
			ID:        "1",
			Firstname: "John",
			Lastname:  "Doe",
			Address: &model.Address{
				City: "City X", State: "State X"}})

	log.Printf(utils.ToString(people))
}

func GetAllPeople() []model.Person {
	return people
}

func FindPersonByID(ID string) model.Person {
	for _, item := range people {
		if item.ID == ID {
			return item
		}
	}
	return model.Person{}
}

func CreatePerson(person model.Person) error {
	people = append(people, person)
	return nil
}

func DeletePerson(person model.Person) error {
	for index, item := range people {
		if item.ID == person.ID {
			people = append(people[:index], people[index+1:]...)
			return nil
		}
	}
	return errors.New("person does not exist")
}
