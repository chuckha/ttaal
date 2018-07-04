package models

import "time"

type Vote struct {
	ID          int64
	StatementID int64
	UserID      int64
	Created     time.Time
	Updated     time.Time
}

func NewVote() *Vote {
	return &Vote{}
}
