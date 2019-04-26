package routehandler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type PersonReqBody struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func init() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{personID}", GetPerson)
	router.Post("/", CreatePerson)
	router.Delete("/{personID}", DeletePerson)
	router.Get("/", GetPeople)
	return router
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personID")
	for _, item := range people {
		if item.ID == personID {
			render.JSON(w, r, item)
		}
	}
	http.Error(w, "Person not found.", http.StatusNotFound)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var personBody PersonReqBody
	err := json.NewDecoder(r.Body).Decode(&personBody)
	if err != nil {
		//TODO: Pack log and error response in a function
		log.Printf("Unable to parse CreatePerson request body: %s\n", err.Error())
		http.Error(w, "Wrong body format or element missing in body.", http.StatusBadRequest)
	}
	newPerson := Person{
		ID:        "123",
		Firstname: personBody.Firstname,
		Lastname:  personBody.Lastname,
		Address:   personBody.Address,
	}
	people = append(people, newPerson)
	response := make(map[string]string)
	response["message"] = "Person created successfully"

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)

	personID := chi.URLParam(r, "personID")
	for index, item := range people {
		if item.ID == personID {
			var person = people[index]
			people = append(people[:index], people[index + 1:]...)
			log.Printf("Deleted person " + toString(person))

			response["message"] = "Person with ID '" + personID + "' deleted successfully"
			log.Printf(response["message"])
			render.JSON(w, r, response)
			return
		}
	}

	response["message"] = "Person with ID: '" + personID + "' not found"
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, response)
}

func toString(v interface{}) string {
	out, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return string(out)
}