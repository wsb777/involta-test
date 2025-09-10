package services

import (
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
)

type DeletePersonService interface {
	DeletePerson(person *dto.PersonDelete) error
}

type deletePersonService struct {
	personRepo repo.ReindexerRepo
}

func NewDeletePersonService(repo repo.ReindexerRepo) DeletePersonService {
	return &deletePersonService{personRepo: repo}
}

func (s *deletePersonService) DeletePerson(person *dto.PersonDelete) error {
	return s.personRepo.DeletePersonByID(person.ID)
}
