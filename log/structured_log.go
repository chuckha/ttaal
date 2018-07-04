package log

import (
	"github.com/sirupsen/logrus"
)

// Structured implements structured logging
type Structured struct{}

// Debugf is analogous to fmt.Printf but with structure
func (s *Structured) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Debugln is analogous ot fmt.Println but with structure
func (s *Structured) Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}
