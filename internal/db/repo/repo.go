package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/restream/reindexer/v5"
	"github.com/wsb777/involta-test/internal/models"
)

type ReindexerRepo interface {
	CreatePerson(ctx context.Context, person *models.Person) error
	GetPersonByID(ctx context.Context, id int) (*models.Person, error)
	UpdatePerson(ctx context.Context, person *models.Person) error
	DeletePersonByID(ctx context.Context, id int) error
	GetPersonsList(ctx context.Context, searchParams *models.SearchParams) ([]*models.Person, error)
}

type reindexerRepo struct {
	db *reindexer.Reindexer
}

func NewReindexerRepo(db *reindexer.Reindexer) ReindexerRepo {
	return &reindexerRepo{db: db}
}

func (r *reindexerRepo) CreatePerson(ctx context.Context, person *models.Person) error {
	tx, err := r.db.WithContext(ctx).BeginTx("persons")
	if err != nil {
		log.Printf("[WARN] Problem with CreatePerson in repo - context timeout. Error message: %s", err)
		return fmt.Errorf("Request timeout")
	}
	tx.Upsert(person, "id=serial()")
	if err := tx.Commit(); err != nil {
		log.Printf("[WARN] Problem with commit person - context timeout. Error message: %s", err)
		return fmt.Errorf("Request timeout")
	}
	return nil
}

func (r *reindexerRepo) GetPersonByID(ctx context.Context, id int) (*models.Person, error) {
	if err := ctx.Err(); err != nil {
		log.Printf("[WARN] Problem with start get person - context timeout. Error message: %s", err)
		return nil, fmt.Errorf("Request timeout")
	}

	query := r.db.Query("persons").Where("id", reindexer.GT, id)
	iterator := query.Exec()
	defer iterator.Close()

	if err := iterator.Error(); err != nil {
		log.Printf("[WARN] Problem with iterator get person - context timeout. Error message: %s", err)
		return nil, fmt.Errorf("Request timeout")
	}

	var data *models.Person

	for iterator.Next() {
		select {
		case <-ctx.Done():
			log.Print("[WARN] Problem with get person - context timeout.")
			return nil, fmt.Errorf("Request timeout")

		default:
			data = iterator.Object().(*models.Person)
		}
	}
	return data, nil
}

func (r *reindexerRepo) UpdatePerson(ctx context.Context, person *models.Person) error {
	tx, err := r.db.WithContext(ctx).BeginTx("persons")
	if err != nil {
		log.Printf("[WARN] Problem with UpdatePerson in repo. Error message: %s", err)
		return fmt.Errorf("Request timeout")
	}
	tx.Update(person)
	if err := tx.Commit(); err != nil {
		log.Printf("[WARN] Problem with commit update person. Error message: %s", err)
		return fmt.Errorf("Request timeout")
	}

	return nil
}

func (r *reindexerRepo) DeletePersonByID(ctx context.Context, id int) error {
	tx, err := r.db.WithContext(ctx).BeginTx("persons")
	if err != nil {
		log.Printf("[WARN] Problem with DeletePerson in repo. Error message: %s", err)
		return fmt.Errorf("Request timeout")
	}
	tx.Query().Where("id", reindexer.EQ, id).Delete()
	if err := tx.Commit(); err != nil {
		log.Printf("[WARN] Problem with commit delete person. Error message: %s", err)
		return fmt.Errorf("Request timeout")
	}
	return nil
}

func (r *reindexerRepo) GetPersonsList(ctx context.Context, searchParams *models.SearchParams) ([]*models.Person, error) {
	var data []*models.Person

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	query := r.db.Query("persons").Match("fullName", "*"+searchParams.Text+"*", "<stemmers>").Where("id", reindexer.GT, searchParams.LastID).Limit(searchParams.Limit)
	iterator := query.Exec()
	defer iterator.Close()

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	for iterator.Next() {
		select {
		case <-ctx.Done():
			log.Print("[WARN] Context timeout")
			return nil, fmt.Errorf("Request timeout")

		default:
			person := iterator.Object().(*models.Person)
			person.Sort = int(iterator.Rank())
			data = append(data, person)
		}
	}

	return data, nil
}
