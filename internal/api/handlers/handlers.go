package handlers

import (
	"net/http"

	"github.com/wsb777/involta-test/internal/api/controllers"
	"github.com/wsb777/involta-test/internal/api/middleware"
)

func NewHTTPServer(createPersonController *controllers.CreatePersonController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/person", createPersonController.CreatePerson)
	middleServer := middleware.AllInfoMiddleware(mux)
	return middleServer
}
