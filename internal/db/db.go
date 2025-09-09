package db

import (
	"fmt"
	"log"

	"github.com/restream/reindexer/v5"
	_ "github.com/restream/reindexer/v5/bindings/cproto"
	"github.com/wsb777/involta-test/internal/config"
	"github.com/wsb777/involta-test/internal/models"
)

func ConnectToDatabase(config *config.Config) *reindexer.Reindexer {
	connectPath := fmt.Sprintf("cproto://%v:%v/%v", config.Host, config.Port, config.DBName)
	log.Print("Connecting to Reindexer...")
	db, err := reindexer.NewReindex(connectPath, reindexer.WithCreateDBIfMissing())

	if err != nil {
		panic(err)
	}
	log.Print("Successful connect!")

	err = db.OpenNamespace("persons", reindexer.DefaultNamespaceOptions(), models.Person{})

	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}
