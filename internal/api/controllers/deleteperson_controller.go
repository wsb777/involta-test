package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wsb777/involta-test/internal/dto"
)

type DeletePersonService interface {
	DeletePerson(person *dto.PersonID) error
}

type DeletePersonController struct {
	service DeletePersonService
}

func NewDeletePersonController(service DeletePersonService) *DeletePersonController {
	return &DeletePersonController{service: service}
}

func (c *DeletePersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person dto.PersonID
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := c.service.DeletePerson(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}
