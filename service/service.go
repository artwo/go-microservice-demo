package service

import "chiapitest/model"

type PeopleService interface {
	GetAllPeople() ([]model.Person, error)
	GetPerson(personID string) (model.Person, error)
	AddPerson(person model.PersonNoID) error
	RemovePerson(personID string) error
}
