package services

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/dto"
	"github.com/wsb777/involta-test/internal/models"
)

type GetPersonsListService interface {
	GetPersonsList(ctx context.Context, searchParamsDTO *dto.SearchParams) ([]dto.PersonGet, error)
}

type getPersonsListService struct {
	personRepo repo.ReindexerRepo
}

type Task struct {
	index int
	value dto.PersonGet
}

func NewGetPersonsListService(repo repo.ReindexerRepo) GetPersonsListService {
	return &getPersonsListService{personRepo: repo}
}

func (s *getPersonsListService) GetPersonsList(ctx context.Context, searchParamsDTO *dto.SearchParams) ([]dto.PersonGet, error) {
	searchParams := &models.SearchParams{
		Text:   searchParamsDTO.Text,
		Offset: searchParamsDTO.Offset,
		Limit:  searchParamsDTO.Limit,
	}
	data, err := s.personRepo.GetPersonsList(ctx, searchParams)
	if err != nil {
		return nil, err
	}

	personsCount := len(data)
	persons := make([]dto.PersonGet, personsCount)

	for i, value := range data {
		documentsCount := len(value.Documents)
		documents := make([]dto.DocumentGet, documentsCount)

		for k, doc := range value.Documents {
			documents[k].ID = doc.ID
			documents[k].Name = doc.Name
			documents[k].CreateAt = doc.CreateAt
		}
		persons[i] = dto.PersonGet{
			ID:         value.ID,
			FirstName:  value.FirstName,
			SecondName: value.SecondName,
			MiddleName: value.MiddleName,
			Sort:       value.Sort,
			CreateAt:   value.CreateAt,
			UpdateAt:   value.UpdateAt,
			Documents:  documents,
		}
	}
	persons, err = s.dataProcessing(ctx, persons)

	if err != nil {
		return nil, err
	}

	return persons, nil
}

/*
# COMMENT: Если эта функция имеет смысл, то лучше её переносить наверх.

Будет два канала - в первый поступает models.Person для обработки, потом передается во второй уже в формате PersonGet.
Так можно одновременно перевести из model.Person в dto.PersonGet и обработать поля.
Сделал отдельной функцией, тк сейчас эта обработка не имеет смысла

Так же по улучшениям:
1) Можно сделать настраиваемое количество воркеров (в зависимости от количества элементов)
2) Можно сделать обработку без каналов, если элементов мало
*/
func (s *getPersonsListService) dataProcessing(ctx context.Context, persons []dto.PersonGet) ([]dto.PersonGet, error) {
	if ctx.Err() != nil {
		return nil, fmt.Errorf("Request timeout")
	}

	var wg sync.WaitGroup

	tasksCh := make(chan Task)
	resultCh := make(chan Task)
	countOfWorkers := 10

	for i := 0; i < countOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, tasksCh, resultCh, &wg)
	}

	go func() {
		defer close(tasksCh)
		for i, person := range persons {
			select {
			case tasksCh <- Task{index: i, value: person}:
			case <-ctx.Done():
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make([]dto.PersonGet, len(persons))
	for res := range resultCh {
		results[res.index] = res.value
	}

	if ctx.Err() != nil {
		return nil, fmt.Errorf("Request timeout")
	}

	return results, nil

}

func worker(ctx context.Context, tasksCh <-chan Task, resultsCh chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasksCh:
			if !ok {
				return
			}
			num := rand.Intn(6)
			numDoc := rand.Intn(2)

			switch num {
			case 0:
				task.value.CreateAt = "Обработан"
			case 1:
				task.value.UpdateAt = "Обработан"
			case 2:
				task.value.FirstName = "Обработан"
			case 3:
				task.value.SecondName = "Обработан"
			case 4:
				task.value.MiddleName = "Обработан"
			}

			for i := range task.value.Documents {
				switch numDoc {
				case 0:
					task.value.Documents[i].Name = "Обработан"
				case 1:
					task.value.Documents[i].Name = "Обработан: " + task.value.Documents[i].Name
				}
			}

			select {
			case resultsCh <- task:
			case <-ctx.Done():
				return
			}
		}
	}
}
