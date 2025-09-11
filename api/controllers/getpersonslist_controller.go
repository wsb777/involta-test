package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonsListService interface {
	GetPersonsList(searhParams *dto.SearchParams) ([]dto.PersonGet, error)
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

	lastId := params.Get("lastId")
	if lastId == "" {
		http.Error(w, "lastId not found", http.StatusBadRequest)
		return
	}

	lastIdNum, err := strconv.Atoi(lastId)
	if err != nil {
		http.Error(w, "lastId not number", http.StatusBadRequest)
		return
	}

	searchParams := &dto.SearchParams{
		Limit:  limitNum,
		LastID: lastIdNum,
		Text:   text,
	}

	if err != nil {
		http.Error(w, "id not a number", http.StatusBadRequest)
		return
	}

	value, err := c.service.GetPersonsList(searchParams)
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
