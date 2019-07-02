package service

import (
	"chiapitest/model"
	"chiapitest/service"
	"chiapitest/test/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedPeople = []model.Person{
	{
		ID:        "2",
		Firstname: "Koko",
		Lastname:  "Doe",
		Address: &model.Address{
			City: "City Z", State: "State Y"},
	},
	{
		ID:        "1",
		Firstname: "John",
		Lastname:  "Doe",
		Address: &model.Address{
			City: "City X", State: "State X"},
	},
}

func TestGetAllPeople(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleRepoMock.On("GetAll").Return(expectedPeople, nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	result, err := peopleService.GetAllPeople()

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "GetAll", 1)
	assert.Equal(t, expectedPeople, result, "People data has changed")
	assert.Nil(t, err, "There is an unexpected error")
}

func TestGetAllPeopleEmpty(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleRepoMock.On("GetAll").Return([]model.Person{}, nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	result, err := peopleService.GetAllPeople()

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "GetAll", 1)
	assert.Equal(t, []model.Person{}, result, "People data is not empty")
	assert.Nil(t, err, "There is an unexpected error")
}

func TestGetPerson(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleRepoMock.On("FindByID", "1").Return(expectedPeople[1], nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	result, err := peopleService.GetPerson("1")

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "FindByID", 1)
	assert.Equal(t, expectedPeople[1], result, "Incorrect person selected")
	assert.Nil(t, err, "There is an unexpected error")
}

func TestGetPersonNotFound(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleRepoMock.On("FindByID", "random-id").Return(model.Person{}, nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	result, err := peopleService.GetPerson("random-id")

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "FindByID", 1)
	assert.Equal(t, model.Person{}, result, "This is not an empty person object")
	assert.Nil(t, err, "There is an unexpected error")
}

func TestRemovePerson(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleRepoMock.On("FindByID", "0").Return(expectedPeople[0], nil)
	peopleRepoMock.On("Remove", expectedPeople[0]).Return(nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	err := peopleService.RemovePerson("0")

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "FindByID", 1)
	peopleRepoMock.AssertNumberOfCalls(t, "Remove", 1)
	assert.Nil(t, err, "There is an unexpected error")
}

func TestRemovePersonNotFound(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleRepoMock.On("FindByID", "random-id").Return(model.Person{}, nil)
	//peopleRepoMock.On("Remove", expectedPeople[0]).Return(nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	err := peopleService.RemovePerson("random-id")

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "FindByID", 1)
	peopleRepoMock.AssertNumberOfCalls(t, "Remove", 0)
	assert.NotNil(t, err, "There was no expected error")
}

func TestAddPerson(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	personNoID := model.PersonNoID{
		Firstname: "John",
		Lastname:  "Wick",
		Address: &model.Address{
			City:  "Merida",
			State: "Merida",
		},
	}
	peopleRepoMock.On(
		"Add",
		model.Person{
			ID:        "123",
			Firstname: personNoID.Firstname,
			Lastname:  personNoID.Lastname,
			Address:   personNoID.Address,
		}).Return(nil)

	peopleService := service.NewPeopleService(peopleRepoMock)
	err := peopleService.AddPerson(personNoID)

	peopleRepoMock.AssertExpectations(t)
	peopleRepoMock.AssertNumberOfCalls(t, "Add", 1)
	assert.Nil(t, err, "There is an unexpected error")
}
