package services

import (
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
	"github.com/wsb777/involta-test/internal/models"
)

type UpdatePersonService interface {
	UpdatePerson(personDto *dto.PersonUpdate) error
}

type updatePersonService struct {
	personRepo repo.ReindexerRepo
}

func NewUpdatePersonService(repo repo.ReindexerRepo) UpdatePersonService {
	return &updatePersonService{personRepo: repo}
}

func (s *updatePersonService) UpdatePerson(personDto *dto.PersonUpdate) error {
	person := &models.Person{
		ID:         personDto.ID,
		FirstName:  personDto.FirstName,
		SecondName: personDto.SecondName,
		MiddleName: personDto.MiddleName,
	}

	return s.personRepo.UpdatePerson(person)
}
