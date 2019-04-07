package routehandler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Employee struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo)
	router.Delete("/{todoID}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Employee {
		ID: todoID,
		Title: "Hello World",
		Body: "Hello world from planet earth",
	}
	render.JSON(w, r, todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response)
}

func CreateTodo(w http.ResponseWriter, r * http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Employee {
		{
			ID: "id",
			Title: "Hello World",
			Body: "Hello world from planet earth",
		},
	}
	render.JSON(w, r, todos)
}