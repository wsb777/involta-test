package controllers

import (
	"net/http"
	"strconv"

	"github.com/wsb777/involta-test/internal/dto"
)

type DeletePersonService interface {
	DeletePerson(person *dto.PersonDelete) error
}

type DeletePersonController struct {
	service DeletePersonService
}

func NewDeletePersonController(service DeletePersonService) *DeletePersonController {
	return &DeletePersonController{service: service}
}

func (c *DeletePersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person dto.PersonDelete

	id := r.PathValue("id")
	idNum, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "id not a number", http.StatusBadRequest)
		return
	}

	person.ID = idNum

	if person.ID == 0 {
		http.Error(w, "request without identifier", http.StatusBadRequest)
		return
	}

	if err := c.service.DeletePerson(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}
