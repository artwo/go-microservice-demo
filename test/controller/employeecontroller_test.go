package controller

import (
	"chiapitest/controller"
	"chiapitest/model"
	"chiapitest/test/mock"
	"chiapitest/utils"
	"context"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
	peopleServiceMock := new(mock.PeopleServiceMock)
	peopleServiceMock.On("GetAllPeople").Return(expectedPeople, nil)
	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.GetAllPeople)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, utils.ToString(expectedPeople)+"\n", rr.Body.String(), "Request returned unexpected response body")
}

func TestGetPerson(t *testing.T) {
	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleServiceMock := new(mock.PeopleServiceMock)
	peopleServiceMock.On("GetPerson", "2").Return(expectedPeople[0], nil)
	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.GetPerson)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	// Set query param
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("personID", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, utils.ToString(expectedPeople[0])+"\n", rr.Body.String(), "Request returned unexpected response body")
}

func TestPostPerson(t *testing.T) {
	var newPerson = model.PersonNoID{
		Firstname: "Test First Name",
		Lastname:  "Test Last Name",
		Address: &model.Address{
			City:  "Test City",
			State: "Test State",
		},
	}
	expectedResponse := make(map[string]string)
	expectedResponse["message"] = "Person created successfully"

	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleServiceMock := new(mock.PeopleServiceMock)
	peopleServiceMock.On("AddPerson", newPerson).Return(nil)
	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.PostPerson)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", utils.ToIoReader(newPerson))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, utils.ToString(expectedResponse)+"\n", rr.Body.String(), "Request returned unexpected response body")
}

func TestDeletePerson(t *testing.T) {
	expectedResponse := make(map[string]string)
	expectedResponse["message"] = "Person deleted successfully"

	peopleRepoMock := new(mock.PeopleRepoMock)
	peopleServiceMock := new(mock.PeopleServiceMock)
	peopleServiceMock.On("RemovePerson", "2").Return(nil)
	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.DeletePerson)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/", nil)

	// Set query param
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("personID", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, utils.ToString(expectedResponse)+"\n", rr.Body.String(), "Request returned unexpected response body")
}
