package repo

import (
	"fmt"
	"log"

	"github.com/restream/reindexer/v5"
	"github.com/wsb777/involta-test/internal/models"
)

type ReindexerRepo interface {
	CreatePerson(person *models.Person) error
	GetPersonByID(id int) (*models.Person, error)
	UpdatePerson(person *models.Person) error
	DeletePersonByID(id int) error
}

type reindexerRepo struct {
	db *reindexer.Reindexer
}

func NewReindexerRepo(db *reindexer.Reindexer) ReindexerRepo {
	return &reindexerRepo{db: db}
}

func (r *reindexerRepo) CreatePerson(person *models.Person) error {
	err := r.db.Upsert("persons", person, "id=serial()")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (r *reindexerRepo) GetPersonByID(id int) (*models.Person, error) {
	data, found := r.db.Query("persons").Where("id", reindexer.EQ, id).Get()

	if found {
		item := data.(*models.Person)
		return item, nil
	}
	return nil, fmt.Errorf("error while get person")
}

func (r *reindexerRepo) UpdatePerson(person *models.Person) error {
	var err error
	countOfUpdate := r.db.Query("persons").Where("id", reindexer.EQ, person.ID).Set("firstName", person.FirstName).Set("secondName", person.SecondName).Set("middleName", person.MiddleName).Set("updateAt", person.UpdateAt).Set("documents.", person.Documents).Update()
	if countOfUpdate.Count() == 0 {
		err = fmt.Errorf("entry is not updated")
	}

	return err
}

func (r *reindexerRepo) DeletePersonByID(id int) error {
	_, err := r.db.Query("persons").Where("id", reindexer.EQ, id).Delete()
	return err
}
