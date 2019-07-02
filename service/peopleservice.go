package service

import (
	"chiapitest/model"
	"chiapitest/repo"
	"errors"
)

type peopleService struct {
	PeopleRepo repo.PeopleRepository
}

func NewPeopleService(peopleRepo repo.PeopleRepository) PeopleService {
	return &peopleService{peopleRepo}
}

func (s *peopleService) GetAllPeople() ([]model.Person, error) {
	return s.PeopleRepo.GetAll()
}

func (s *peopleService) GetPerson(personID string) (model.Person, error) {
	return s.PeopleRepo.FindByID(personID)
}

func (s *peopleService) AddPerson(person model.PersonNoID) error {
	newPerson := model.Person{
		ID:        "123",
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Address:   person.Address,
	}
	if err := s.PeopleRepo.Add(newPerson); err != nil {
		return errors.New("Unable to create person error: '" + err.Error())
	}
	return nil
}

func (s *peopleService) RemovePerson(personID string) error {
	person, err := s.PeopleRepo.FindByID(personID)
	if err != nil {
		return err
	}
	if (person == model.Person{}) {
		return errors.New("Unable to find person with ID '" + personID + "'")
	}

	return s.PeopleRepo.Remove(person)
}
