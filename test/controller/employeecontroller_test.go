package controller

import (
	"bytes"
	"chiapitest/controller"
	"chiapitest/model"
	m "chiapitest/test/mock"
	"chiapitest/utils"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

var expectedPeople = []model.Person{
	{
		ID:        "2",
		FirstName: "Koko",
		LastName:  "Doe",
		Address: &model.Address{
			City: "City Z", State: "State Y"},
	}, {
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Address: &model.Address{
			City: "City X", State: "State X"},
	},
}

func TestGetAllPeople(t *testing.T) {
	peopleRepoMock := new(m.PeopleRepoMock)
	peopleServiceMock := new(m.PeopleServiceMock)
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
	peopleRepoMock := new(m.PeopleRepoMock)
	peopleServiceMock := new(m.PeopleServiceMock)
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
		FirstName: "Test First Name",
		LastName:  "Test Last Name",
		Address: &model.Address{
			City:  "Test City",
			State: "Test State",
		},
	}
	expectedResponse := make(map[string]string)
	expectedResponse["message"] = "Person created successfully"

	peopleRepoMock := new(m.PeopleRepoMock)
	peopleServiceMock := new(m.PeopleServiceMock)
	peopleServiceMock.On("AddPerson", newPerson).Return(nil)
	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.PostPerson)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", utils.ToIoReader(newPerson))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, utils.ToString(expectedResponse)+"\n", rr.Body.String(), "Request returned unexpected response body")
}

func postWithoutPersonParam(p string) *httptest.ResponseRecorder {
	requestBodyMap := map[string]interface{}{
		"firstName": "Test Last Name",
		"lastName":  "Test Last Name",
		"address": map[string]string{
			"city":  "Test City",
			"state": "Test State",
		},
	}
	delete(requestBodyMap, p)
	return unmockedPostRequest(requestBodyMap)
}

func postPersonWithoutAddressParam(p string) *httptest.ResponseRecorder {
	address := map[string]string{
		"city":  "Test City",
		"state": "Test State",
	}
	delete(address, p)
	requestBodyMap := map[string]interface{}{
		"firstName": "Test Last Name",
		"lastName":  "Test Last Name",
		"address":   address,
	}
	return unmockedPostRequest(requestBodyMap)
}

func unmockedPostRequest(requestBodyMap map[string]interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(requestBodyMap)

	peopleRepoMock := new(m.PeopleRepoMock)
	peopleServiceMock := new(m.PeopleServiceMock)

	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.PostPerson)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(requestBody))

	handler.ServeHTTP(rr, req)
	return rr
}

func TestPostPersonWithoutFirstName(t *testing.T) {
	expectedResponse := "Wrong body format or element missing in body\n"
	rr := postWithoutPersonParam("firstName")

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, expectedResponse, rr.Body.String(), "Request returned unexpected response body")
}

func TestPostPersonWithoutLastName(t *testing.T) {
	expectedResponse := "Wrong body format or element missing in body\n"
	rr := postWithoutPersonParam("lastName")

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, expectedResponse, rr.Body.String(), "Request returned unexpected response body")
}

func TestPostPersonWithoutAddress(t *testing.T) {
	expectedResponse := "Wrong body format or element missing in body\n"
	rr := postWithoutPersonParam("address")

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, expectedResponse, rr.Body.String(), "Request returned unexpected response body")
}

func TestPostPersonWithoutCity(t *testing.T) {
	expectedResponse := "Wrong body format or element missing in body\n"
	rr := postPersonWithoutAddressParam("city")

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, expectedResponse, rr.Body.String(), "Request returned unexpected response body")
}

func TestPostPersonWithoutState(t *testing.T) {
	expectedResponse := "Wrong body format or element missing in body\n"
	rr := postPersonWithoutAddressParam("state")

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, expectedResponse, rr.Body.String(), "Request returned unexpected response body")
}

func TestDeletePerson(t *testing.T) {
	expectedResponse := make(map[string]string)
	expectedResponse["message"] = "Person deleted successfully"

	peopleRepoMock := new(m.PeopleRepoMock)
	peopleServiceMock := new(m.PeopleServiceMock)
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

func TestDeletePersonWithoutID(t *testing.T) {
	expectedResponse := "Request path parameter personID is missing\n"

	peopleRepoMock := new(m.PeopleRepoMock)
	peopleServiceMock := new(m.PeopleServiceMock)
	peopleController := controller.NewTestController(peopleRepoMock, peopleServiceMock)

	handler := http.HandlerFunc(peopleController.DeletePerson)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/", nil)

	// Set query param
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request returned unexpected status code")
	assert.Equal(t, expectedResponse, rr.Body.String(), "Request returned unexpected response body")
}
