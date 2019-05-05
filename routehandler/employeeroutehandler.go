package routehandler

import (
	"chiapitest/models"
	"chiapitest/service"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getPeople)
	router.Get("/{personID}", getPerson)
	router.Post("/", createPerson)
	router.Delete("/{personID}", deletePerson)
	return router
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	people := service.GetPeople()
	render.JSON(w, r, people)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personID")
	person := service.GetPerson(personID)
	if (person == models.Person{}) {
		log.Printf("Unable to find person with ID: %s\n", personID)
		http.Error(w, "Person not found.", http.StatusNotFound)
		return
	}
	render.JSON(w, r, person)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var personRequest models.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&personRequest)
	if err != nil {
		//TODO: Pack log and error response in a function
		log.Printf("Unable to parse CreatePerson request body: %s\n", err.Error())
		http.Error(w, "Wrong body format or element missing in body.", http.StatusBadRequest)
		return
	}

	err = service.AddPerson(personRequest)
	if err != nil {
		log.Printf("Unable to store new person: %s\n", err.Error())
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	response := make(map[string]string)
	response["message"] = "Person created successfully"
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personID")
	response := make(map[string]string)

	err := service.RemovePerson(personID)
	if err != nil {
		log.Printf("Unable to find person with ID: %s\n", personID)
		//TODO: render Errors with renderer
		http.Error(w, "Person not found.", http.StatusNotFound)
		return
	}

	log.Printf("Deleted person " + personID)
	response["message"] = "Person deleted successfully."
	render.JSON(w, r, response)
}