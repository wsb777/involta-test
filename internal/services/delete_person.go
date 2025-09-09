package services

import (
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
)

type DeletePersonService interface {
	DeletePerson(person *dto.PersonID) error
}

type deletePersonService struct {
	personRepo repo.ReindexerRepo
}

func NewDeletePersonService(repo repo.ReindexerRepo) DeletePersonService {
	return &deletePersonService{personRepo: repo}
}

func (s *deletePersonService) DeletePerson(person *dto.PersonID) error {
	return s.personRepo.DeletePersonByID(person.ID)
}
