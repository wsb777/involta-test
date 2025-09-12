package services

import (
	"context"
	"time"

	"github.com/wsb777/involta-test/internal/cache"
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
	"github.com/wsb777/involta-test/internal/models"
)

type UpdatePersonService interface {
	UpdatePerson(ctx context.Context, personDto *dto.PersonUpdate) error
}

type updatePersonService struct {
	personRepo repo.ReindexerRepo
	memStore   *cache.MemStore
}

func NewUpdatePersonService(repo repo.ReindexerRepo, personCache *cache.MemStore) UpdatePersonService {
	return &updatePersonService{personRepo: repo, memStore: personCache}
}

func (s *updatePersonService) UpdatePerson(ctx context.Context, personDto *dto.PersonUpdate) error {

	timeNow := time.Now().String()

	documentsCount := len(personDto.Documents)
	documents := make([]models.Document, documentsCount)

	for i, doc := range personDto.Documents {
		documents[i].ID = doc.ID
		documents[i].Name = doc.Name
		documents[i].CreateAt = timeNow
	}

	person := &models.Person{
		ID:         personDto.ID,
		FirstName:  personDto.FirstName,
		SecondName: personDto.SecondName,
		MiddleName: personDto.MiddleName,
		Documents:  documents,
		UpdateAt:   timeNow,
	}

	err := s.personRepo.UpdatePerson(ctx, person)

	if err == nil {
		s.memStore.Set(person.ID, person)
	}

	return err
}
