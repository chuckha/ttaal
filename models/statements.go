package models

import (
	"time"
)

const (
	// StatementsTable is the name of the statements table
	StatementsTable = "statements"
)

// Statement is a statement in two truths and a lie
type Statement struct {
	ID        int64     `sql:"id,primary,autoinc"`
	PollID    int64     `sql:"poll_id"`
	IsALie    bool      `sql:"is_a_lie"`
	Statement string    `sql:"statement"`
	Created   time.Time `sql:"created"`
	Updated   time.Time `sql:"updated"`
}

// NewStatement returns an instantiated statement
func NewStatement(statement string, pollID int64, isALie bool) *Statement {
	return &Statement{
		PollID:    pollID,
		IsALie:    isALie,
		Statement: statement,
	}
}
