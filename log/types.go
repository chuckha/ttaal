package log

// Logger defines the usable logging methods.
type Logger interface {
	// Debugf is analogous to fmt.Printf
	Debugf(string, ...interface{})
	// Debugln is analogous ot fmt.Println
	Debugln(...interface{})
	// Infoln is analogous to fmt.Println
	Infoln(...interface{})
	// Infof is analogous to fmt.Printf
	Infof(string, ...interface{})
}
