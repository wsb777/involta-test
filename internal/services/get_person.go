package services

import (
	"context"

	"github.com/wsb777/involta-test/internal/cache"
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonService interface {
	GetPerson(ctx context.Context, person *dto.PersonGet) (*dto.PersonGet, error)
}

type getPersonService struct {
	personRepo repo.ReindexerRepo
	memStore   *cache.MemStore
}

func NewGetPersonService(repo repo.ReindexerRepo, personCache *cache.MemStore) GetPersonService {
	return &getPersonService{personRepo: repo, memStore: personCache}
}

func (s *getPersonService) GetPerson(ctx context.Context, person *dto.PersonGet) (*dto.PersonGet, error) {
	var err error

	personModel, exist := s.memStore.Get(person.ID)
	if !exist {
		personModel, err = s.personRepo.GetPersonByID(ctx, person.ID)
		if err != nil {
			return nil, err
		}
	}

	documentsCount := len(personModel.Documents)
	documents := make([]dto.DocumentGet, documentsCount)

	for i, doc := range personModel.Documents {
		documents[i].ID = doc.ID
		documents[i].Name = doc.Name
		documents[i].CreateAt = doc.CreateAt
	}

	personDto := &dto.PersonGet{
		ID:         personModel.ID,
		FirstName:  personModel.FirstName,
		SecondName: personModel.SecondName,
		MiddleName: personModel.MiddleName,
		Documents:  documents,
		CreateAt:   personModel.CreateAt,
		UpdateAt:   personModel.UpdateAt,
	}

	s.memStore.Set(person.ID, personModel)

	return personDto, err
}
