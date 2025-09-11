package services

import (
	"time"

	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
	"github.com/wsb777/involta-test/internal/models"
)

type CreatePersonService interface {
	CreatePerson(personDto *dto.PersonCreate) error
}

type createPersonService struct {
	personRepo repo.ReindexerRepo
}

func NewCreatePersonService(repo repo.ReindexerRepo) CreatePersonService {
	return &createPersonService{personRepo: repo}
}

func (s *createPersonService) CreatePerson(personDto *dto.PersonCreate) error {
	timeNow := time.Now().String()

	documentsCount := len(personDto.Documents)
	documents := make([]models.Document, documentsCount)

	for i, doc := range personDto.Documents {
		documents[i].ID = doc.ID
		documents[i].Name = doc.Name
		documents[i].CreateAt = timeNow
	}
	person := &models.Person{
		FirstName:  personDto.FirstName,
		SecondName: personDto.SecondName,
		MiddleName: personDto.MiddleName,
		Documents:  documents,
		CreateAt:   timeNow,
		UpdateAt:   timeNow,
	}

	return s.personRepo.CreatePerson(person)
}
