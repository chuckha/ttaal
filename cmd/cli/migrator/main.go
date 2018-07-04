package main

import (
	"fmt"

	"github.com/chuckha/ttaal/data/connections/mysql"
	"github.com/chuckha/ttaal/log"
	"github.com/chuckha/ttaal/models"
)

const (
	database = "ttaal"
)

func main() {
	var logger log.Logger
	logger = &log.Log{Debug: true}

	migrator := mysql.Migrations(logger)

	tables := []string{
		models.StatementsTable,
		models.PollsTable,
	}

	for _, table := range tables {
		if migrator.TableExists(database, table) {
			continue
		}

		if err := migrator.CreateTable(table, tableToModel(table)); err != nil {
			logger.Infof("encountered an error creating table %q: %v\n", table, err)
			continue
		}
	}

	// connect to db
	// check all tables exist
	// if they don't exist create them
	// if they do exist
	// ensure they are what we expect
	// if they are not
	// figure out the diff and apply the update
	// if they are what we expect
	// good.
}

func tableToModel(table string) interface{} {
	switch table {
	case models.StatementsTable:
		return &models.Statement{}
	case models.PollsTable:
		return &models.Poll{}
	default:
		panic(fmt.Sprintf("unknown table %v", table))
	}
}
