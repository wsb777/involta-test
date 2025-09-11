package handlers

import (
	"net/http"

	"github.com/wsb777/involta-test/api/controllers"
	"github.com/wsb777/involta-test/api/middleware"
)

func NewHTTPServer(createPersonController *controllers.CreatePersonController,
	deletePersonController *controllers.DeletePersonController,
	updatePersonController *controllers.UpdatePersonController,
	getPersonController *controllers.GetPersonController,
	getPersonsListController *controllers.GetPersonsListController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/person", createPersonController.CreatePerson)
	mux.HandleFunc("DELETE /api/v1/person/{id}", deletePersonController.DeletePerson)
	mux.HandleFunc("PUT /api/v1/person/{id}", updatePersonController.UpdatePerson)
	mux.HandleFunc("GET /api/v1/person/{id}", getPersonController.GetPerson)
	mux.HandleFunc("GET /api/v1/persons", getPersonsListController.GetPersonsList)
	server := middleware.LogRequestInfoMiddleware(mux)
	return server
}
