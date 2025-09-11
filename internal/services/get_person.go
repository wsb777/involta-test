package services

import (
	"github.com/wsb777/involta-test/internal/cache"
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonService interface {
	GetPerson(person *dto.PersonGet) (*dto.PersonGet, error)
}

type getPersonService struct {
	personRepo repo.ReindexerRepo
	memStore   *cache.MemStore
}

func NewGetPersonService(repo repo.ReindexerRepo, personCache *cache.MemStore) GetPersonService {
	return &getPersonService{personRepo: repo, memStore: personCache}
}

func (s *getPersonService) GetPerson(person *dto.PersonGet) (*dto.PersonGet, error) {
	var err error

	personModel, exist := s.memStore.Get(person.ID)
	if !exist {
		personModel, err = s.personRepo.GetPersonByID(person.ID)
		if err != nil {
			return nil, err
		}
	}

	personDto := &dto.PersonGet{
		ID:         personModel.ID,
		FirstName:  personModel.FirstName,
		SecondName: personModel.SecondName,
		MiddleName: personModel.MiddleName,
	}

	s.memStore.Set(person.ID, personModel)

	return personDto, err
}
