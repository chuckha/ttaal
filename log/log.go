package log

import "fmt"

// Log implements the Logger interface
type Log struct {
	Debug bool
}

// Debugf is analogous to fmt.Printf
func (l *Log) Debugf(format string, args ...interface{}) {
	if l.Debug {
		fmt.Printf(format, args...)
	}
}

// Debugln is analogous ot fmt.Println
func (l *Log) Debugln(args ...interface{}) {
	if l.Debug {
		fmt.Println(args...)
	}
}

// Infoln is analogous ot fmt.Println
func (l *Log) Infoln(args ...interface{}) {
	fmt.Println(args...)
}

// Infof is analogous to fmt.Printf
func (l *Log) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
