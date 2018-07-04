package models

import (
	"time"
)

const (
	PollsTable = "polls"
	dateFormat = "2006-01-02 15:04:05"
)

type Poll struct {
	ID       int64     `sql:"id"`
	UserID   string    `sql:"user_id"`
	Ended    bool      `sql:"ended"`
	Preamble string    `sql:"preamble"`
	Created  time.Time `sql:"created"`
	Updated  time.Time `sql:"updated"`
}

func NewPoll() *Poll {
	t := time.Now()
	return &Poll{
		Created: t,
		Updated: t,
	}
}

// func CreatePollTable(db Execer) error {
// 	_, err := db.Exec(`CREATE TABLE polls (
// 		id INT PRIMARY KEY AUTO_INCREMENT,
// 		user_id VARCHAR(100),
// 		ended TINYINT,
// 		preamble TEXT,
// 		created TIMESTAMP,
// 		updated TIMESTAMP
// 		)`)
// 	return err
// }
