package services

import (
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
	"github.com/wsb777/involta-test/internal/models"
)

type GetPersonsListService interface {
	GetPersonsList(searchParamsDTO *dto.SearchParams) ([]dto.PersonGet, error)
}

type getPersonsListService struct {
	personRepo repo.ReindexerRepo
}

func NewGetPersonsListService(repo repo.ReindexerRepo) GetPersonsListService {
	return &getPersonsListService{personRepo: repo}
}

func (s *getPersonsListService) GetPersonsList(searchParamsDTO *dto.SearchParams) ([]dto.PersonGet, error) {
	searchParams := &models.SearchParams{
		Text:   searchParamsDTO.Text,
		LastID: searchParamsDTO.LastID,
		Limit:  searchParamsDTO.Limit,
	}
	data, err := s.personRepo.GetPersonsList(searchParams)
	if err != nil {
		return nil, err
	}

	personsCount := len(data)
	persons := make([]dto.PersonGet, personsCount)

	for i, value := range data {
		documentsCount := len(value.Documents)
		documents := make([]dto.DocumentGet, documentsCount)

		for k, doc := range value.Documents {
			documents[k].ID = doc.ID
			documents[k].Name = doc.Name
			documents[k].CreateAt = doc.CreateAt
		}
		persons[i] = dto.PersonGet{
			ID:         value.ID,
			FirstName:  value.FirstName,
			SecondName: value.SecondName,
			MiddleName: value.MiddleName,
			Sort:       value.Sort,
			CreateAt:   value.CreateAt,
			UpdateAt:   value.UpdateAt,
			Documents:  documents,
		}
	}

	return persons, nil
}
