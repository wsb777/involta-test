package services

import (
	"time"

	"github.com/wsb777/involta-test/internal/cache"
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
	"github.com/wsb777/involta-test/internal/models"
)

type UpdatePersonService interface {
	UpdatePerson(personDto *dto.PersonUpdate) error
}

type updatePersonService struct {
	personRepo repo.ReindexerRepo
	memStore   *cache.MemStore
}

func NewUpdatePersonService(repo repo.ReindexerRepo, personCache *cache.MemStore) UpdatePersonService {
	return &updatePersonService{personRepo: repo, memStore: personCache}
}

func (s *updatePersonService) UpdatePerson(personDto *dto.PersonUpdate) error {

	person := &models.Person{
		ID:         personDto.ID,
		FirstName:  personDto.FirstName,
		SecondName: personDto.SecondName,
		MiddleName: personDto.MiddleName,
		UpdateAt:   time.Now().String(),
	}

	err := s.personRepo.UpdatePerson(person)

	if err == nil {
		s.memStore.Set(person.ID, person)
	}

	return err
}
