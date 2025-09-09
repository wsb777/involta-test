package services

import (
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
	person := &models.Person{
		FirstName:  personDto.FirstName,
		SecondName: personDto.SecondName,
		MiddleName: personDto.MiddleName,
	}

	return s.personRepo.CreatePerson(person)
}
