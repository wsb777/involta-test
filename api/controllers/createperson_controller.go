package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wsb777/involta-test/internal/dto"
)

type CreatePersonService interface {
	CreatePerson(ctx context.Context, person *dto.PersonCreate) error
}

type CreatePersonController struct {
	service CreatePersonService
}

func NewCreatePersonController(service CreatePersonService) *CreatePersonController {
	return &CreatePersonController{service: service}
}

func (c *CreatePersonController) CreatePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person dto.PersonCreate

	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := c.service.CreatePerson(ctx, &person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}
