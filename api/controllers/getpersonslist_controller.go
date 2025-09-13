package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonsListService interface {
	GetPersonsList(ctx context.Context, searhParams *dto.SearchParams) ([]dto.PersonGet, error)
}

type GetPersonsListController struct {
	service GetPersonsListService
}

func NewGetPersonsListController(service GetPersonsListService) *GetPersonsListController {
	return &GetPersonsListController{service: service}
}

func (c *GetPersonsListController) GetPersonsList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()

	ctx := r.Context()

	limit := params.Get("limit")
	if limit == "" {
		http.Error(w, "limit not found", http.StatusBadRequest)
		return
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "limit not number", http.StatusBadRequest)
		return
	}

	text := params.Get("text")
	if text == "" {
		http.Error(w, "text not found", http.StatusBadRequest)
		return
	}

	offset := params.Get("offset")
	if offset == "" {
		http.Error(w, "offset not found", http.StatusBadRequest)
		return
	}

	offsetNum, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "offset not number", http.StatusBadRequest)
		return
	}

	searchParams := &dto.SearchParams{
		Limit:  limitNum,
		Offset: offsetNum,
		Text:   text,
	}

	value, err := c.service.GetPersonsList(ctx, searchParams)
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
