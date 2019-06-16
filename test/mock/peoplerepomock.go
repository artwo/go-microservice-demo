package mock

import (
	"chiapitest/model"
	"github.com/stretchr/testify/mock"
)

type PeopleRepoMock struct {
	mock.Mock
}

func (m *PeopleRepoMock) GetAll() ([]model.Person, error) {
	args := m.Called()
	return args.Get(0).([]model.Person), args.Error(1)
}

func (m *PeopleRepoMock) FindByID(ID string) (model.Person, error) {
	args := m.Called(ID)
	return args.Get(0).(model.Person), args.Error(1)
}

func (m *PeopleRepoMock) Add(person model.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *PeopleRepoMock) Remove(person model.Person) error {
	args := m.Called(person)
	return args.Error(0)
}
