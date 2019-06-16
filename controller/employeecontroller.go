package controller

import (
	"chiapitest/model"
	"chiapitest/repo"
	"chiapitest/service"
	"chiapitest/utils"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type Controller struct {
	PeopleRepo    repo.PeopleRepository
	PeopleService service.PeopleService
}

func NewController() *Controller {
	var peopleRepo = repo.NewInMemoryPeopleRepository()
	log.Printf(utils.ToString(peopleRepo))

	return &Controller{
		peopleRepo,
		service.NewPeopleService(peopleRepo),
	}
}

func (c *Controller) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", c.getAllPeople)
	router.Get("/{personID}", c.getPerson)
	router.Post("/", c.postPerson)
	router.Delete("/{personID}", c.deletePerson)
	return router
}

func (c *Controller) getAllPeople(w http.ResponseWriter, r *http.Request) {
	people, err := c.PeopleService.GetAllPeople()
	if err != nil {
		log.Printf("Unable to get all people: %s\n", err.Error())
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, people)
}

func (c *Controller) getPerson(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personID")
	person, err := c.PeopleService.GetPerson(personID)
	if err != nil {
		log.Printf("Unable to get person: %s\n", err.Error())
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	if (person == model.Person{}) {
		log.Printf("Unable to find person with ID: %s\n", personID)
		http.Error(w, "Person not found.", http.StatusNotFound)
		return
	}
	render.JSON(w, r, person)
}

func (c *Controller) postPerson(w http.ResponseWriter, r *http.Request) {
	var personRequest model.PersonNoID
	err := json.NewDecoder(r.Body).Decode(&personRequest)
	if err != nil {
		log.Printf("Unable to parse CreatePerson request body: %s\n", err.Error())
		http.Error(w, "Wrong body format or element missing in body.", http.StatusBadRequest)
		return
	}

	err = c.PeopleService.AddPerson(personRequest)
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

func (c *Controller) deletePerson(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personID")
	response := make(map[string]string)

	err := c.PeopleService.RemovePerson(personID)
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
