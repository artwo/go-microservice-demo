package repo

import "chiapitest/model"

type PeopleRepository interface {
	GetAll() ([]model.Person, error)
	FindByID(ID string) (model.Person, error)
	Add(person model.Person) error
	Remove(person model.Person) error
}
