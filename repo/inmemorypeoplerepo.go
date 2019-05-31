package repo

import (
	"chiapitest/model"
	"errors"
)

type inMemoryPeopleRepository struct {
	People []model.Person
}

func NewInMemoryPeopleRepository() PeopleRepository {
	var people []model.Person
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

	return &inMemoryPeopleRepository{people}
}

//func init() {
//
//	log.Printf(utils.ToString(people))
//}

func (i *inMemoryPeopleRepository) GetAll() ([]model.Person, error) {
	return i.People, nil
}

func (i *inMemoryPeopleRepository) FindByID(ID string) (model.Person, error) {
	for _, item := range i.People {
		if item.ID == ID {
			return item, nil
		}
	}
	return model.Person{}, nil
}

func (i *inMemoryPeopleRepository) Add(person model.Person) error {
	i.People = append(i.People, person)
	return nil
}

func (i *inMemoryPeopleRepository) Remove(person model.Person) error {
	for index, item := range i.People {
		if item.ID == person.ID {
			i.People = append(i.People[:index], i.People[index+1:]...)
			return nil
		}
	}
	return errors.New("person does not exist")
}
