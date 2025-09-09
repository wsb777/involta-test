package controllers

import (
	"encoding/json"
	"net/http"

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
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

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
