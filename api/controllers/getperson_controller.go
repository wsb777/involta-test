package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonService interface {
	GetPerson(person *dto.PersonGet) (*dto.PersonGet, error)
}

type GetPersonController struct {
	service GetPersonService
}

func NewGetPersonController(service GetPersonService) *GetPersonController {
	return &GetPersonController{service: service}
}

func (c *GetPersonController) GetPerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person dto.PersonGet
	id := r.PathValue("id")
	idNum, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "id not a number", http.StatusBadRequest)
		return
	}

	person.ID = idNum

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
