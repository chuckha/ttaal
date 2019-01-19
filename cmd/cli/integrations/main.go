package main

import (
	"github.com/chuckha/modeler/connections/mysql"
	"github.com/chuckha/ttaal/log"
	"github.com/chuckha/ttaal/models"
)

const (
	database = "ttaal"
)

func main() {
	logger := &log.Log{}
	storage := mysql.Storage(logger)

	statement := models.NewStatement("i am a walrus", 0, true)
	id, err := storage.Create(models.NewData(statement, models.StatementsTable))
	if err != nil {
		panic(err)
	}
	logger.Infoln("Created statement with id", id)
}
