package app

import (
	"net/http"
	"time"

	"github.com/wsb777/involta-test/api/controllers"
	"github.com/wsb777/involta-test/api/handlers"
	"github.com/wsb777/involta-test/internal/cache"
	"github.com/wsb777/involta-test/internal/config"
	"github.com/wsb777/involta-test/internal/db"
	"github.com/wsb777/involta-test/internal/db/repo"
	"github.com/wsb777/involta-test/internal/services"
)

func StartApp() (http.Handler, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// База данных
	database := db.ConnectToDatabase(cfg)

	// Репозиторий
	repository := repo.NewReindexerRepo(database)

	// MemStore

	memStore := cache.NewMemStore(time.Minute)
	defer memStore.Stop()

	// Сервисы
	createPersonService := services.NewCreatePersonService(repository)
	deletePersonService := services.NewDeletePersonService(repository, memStore)
	updatePersonService := services.NewUpdatePersonService(repository, memStore)
	getPersonService := services.NewGetPersonService(repository, memStore)

	// Контроллеры
	createPersonController := controllers.NewCreatePersonController(createPersonService)
	deletePersonController := controllers.NewDeletePersonController(deletePersonService)
	updatePersonController := controllers.NewUpdatePersonController(updatePersonService)
	getPersonController := controllers.NewGetPersonController(getPersonService)

	server := handlers.NewHTTPServer(createPersonController, deletePersonController, updatePersonController, getPersonController)
	return server, err
}
