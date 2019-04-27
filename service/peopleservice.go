package service

import (
	"encoding/json"
	"errors"
	"log"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type PersonReqBody struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func init() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	log.Printf(ToString(people))
}

func GetPeople() []Person {
	return people
}

func GetPerson(personID string) Person {
	for _, item := range people {
		if item.ID == personID {
			return item
		}
	}
	return Person{}
}

func CreatePerson(newperson PersonReqBody) error {
	newPerson := Person{
		ID:        "123",
		Firstname: newperson.Firstname,
		Lastname:  newperson.Lastname,
		Address:   newperson.Address,
	}
	people = append(people, newPerson)
	return nil
}

func DeletePerson(personID string) error {
	for index, item := range people {
		if item.ID == personID {
			var person = people[index]
			people = append(people[:index], people[index + 1:]...)

			log.Printf("Deleted person " + ToString(person))
			return nil
		}
	}
	return errors.New("Unable to find person with ID '" + personID + "'")
}

func ToString(v interface{}) string {
	out, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return string(out)
}
