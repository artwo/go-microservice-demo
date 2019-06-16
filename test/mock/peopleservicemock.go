package mock

import (
	"chiapitest/model"
	"github.com/stretchr/testify/mock"
)

type PeopleServiceMock struct {
	mock.Mock
}

func (m *PeopleServiceMock) GetAllPeople() ([]model.Person, error) {
	args := m.Called()
	return args.Get(0).([]model.Person), args.Error(1)
}

func (m *PeopleServiceMock) GetPerson(personID string) (model.Person, error) {
	args := m.Called(personID)
	return args.Get(0).(model.Person), args.Error(1)
}

func (m *PeopleServiceMock) AddPerson(person model.PersonNoID) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *PeopleServiceMock) RemovePerson(personID string) error {
	args := m.Called(personID)
	return args.Error(0)
}
