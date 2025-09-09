package handlers

import (
	"net/http"

	"github.com/wsb777/involta-test/internal/api/controllers"
	"github.com/wsb777/involta-test/internal/api/middleware"
)

func NewHTTPServer(createPersonController *controllers.CreatePersonController,
	deletePersonController *controllers.DeletePersonController,
	updatePersonController *controllers.UpdatePersonController,
	getPersonController *controllers.GetPersonController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /person", createPersonController.CreatePerson)
	mux.HandleFunc("DELETE /person", deletePersonController.DeletePerson)
	mux.HandleFunc("PUT /person", updatePersonController.UpdatePerson)
	mux.HandleFunc("GET /person", getPersonController.GetPerson)
	server := middleware.AllInfoMiddleware(mux)
	return server
}
