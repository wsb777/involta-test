package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonService interface {
	GetPerson(person *dto.PersonID) (*dto.PersonGet, error)
}

type GetPersonController struct {
	service GetPersonService
}

func NewGetPersonController(service GetPersonService) *GetPersonController {
	return &GetPersonController{service: service}
}

func (c *GetPersonController) GetPerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person dto.PersonID
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	value, err := c.service.GetPerson(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonData))
}
