package services

import (
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
)

type GetPersonService interface {
	GetPerson(person *dto.PersonGet) (*dto.PersonGet, error)
}

type getPersonService struct {
	personRepo repo.ReindexerRepo
}

func NewGetPersonService(repo repo.ReindexerRepo) GetPersonService {
	return &getPersonService{personRepo: repo}
}

func (s *getPersonService) GetPerson(person *dto.PersonGet) (*dto.PersonGet, error) {
	personModel, err := s.personRepo.GetPersonByID(person.ID)

	if err != nil {
		return nil, err
	}
	personDto := &dto.PersonGet{
		ID:         personModel.ID,
		FirstName:  personModel.FirstName,
		SecondName: personModel.SecondName,
		MiddleName: personModel.MiddleName,
	}
	return personDto, err
}
