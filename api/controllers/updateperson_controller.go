package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wsb777/involta-test/internal/dto"
)

type UpdatePersonService interface {
	UpdatePerson(person *dto.PersonUpdate) error
}

type UpdatePersonController struct {
	service UpdatePersonService
}

func NewUpdatePersonController(service UpdatePersonService) *UpdatePersonController {
	return &UpdatePersonController{service: service}
}

func (c *UpdatePersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person dto.PersonUpdate
	id := r.PathValue("id")
	idNum, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "id not a number", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	person.ID = idNum

	if person.ID == 0 {
		http.Error(w, "request without identifier", http.StatusBadRequest)
		return
	}

	if err := c.service.UpdatePerson(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}
