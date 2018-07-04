package models_test

import (
	"testing"
	"time"

	"github.com/chuckha/ttaal/models"
)

type s struct {
	Exported   string    `sql:"exported"`
	unexported string    `sql:"unexported"`
	Various    time.Time `sql:"various"`
	Other      int64     `sql:"other"`
	Things     []byte    `sql:"things"`
	NonSql     []string
}

func TestNewData(t *testing.T) {
	test := &s{
		Exported:   "hello",
		unexported: "world",
		Various:    time.Now(),
		Other:      int64(10000),
		Things:     []byte("gopher"),
	}
	m := models.NewData(test, "tablename")
	if len(m.Fields()) != 4 {
		t.Fatalf("expected 4 but got %v", len(m.Fields()))
	}
	if len(m.Values()) != 4 {
		t.Fatalf("expected 4 but got %v", len(m.Values()))
	}
	if m.Table() != "tablename" {
		t.Fatalf("expected 'tablename' but got %v", m.Table())
	}
}
