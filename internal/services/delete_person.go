package services

import (
	"github.com/wsb777/involta-test/internal/cache"
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
)

type DeletePersonService interface {
	DeletePerson(person *dto.PersonDelete) error
}

type deletePersonService struct {
	personRepo repo.ReindexerRepo
	memStore   *cache.MemStore
}

func NewDeletePersonService(repo repo.ReindexerRepo, personCache *cache.MemStore) DeletePersonService {
	return &deletePersonService{personRepo: repo, memStore: personCache}
}

func (s *deletePersonService) DeletePerson(person *dto.PersonDelete) error {
	err := s.personRepo.DeletePersonByID(person.ID)
	if err == nil {
		s.memStore.Delete(person.ID)
	}

	return err
}
