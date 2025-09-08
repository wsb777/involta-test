package app

import (
	"net/http"

	"github.com/wsb777/involta-test/internal/api/controllers"
	"github.com/wsb777/involta-test/internal/api/handlers"
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
	database := db.ConnectToDatabase(cfg)

	repository := repo.NewReindexerRepo(database)

	createPersonService := services.NewCreatePersonService(repository)
	createPersonController := controllers.NewCreatePersonController(createPersonService)
	server := handlers.NewHTTPServer(createPersonController)
	return server, err
}
