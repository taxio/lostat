package log

import (
	"io"
	"log"
)

type dummy struct{}

func (d *dummy) Write(p []byte) (n int, err error) {
	return 0, nil
}

var logger = log.New(&dummy{}, "", 0)

// SetVerboseLogger configure logger for verbose mode
func SetVerboseLogger(out io.Writer) {
	logger = log.New(out, "[DEBUG] ", 0)
}

// Printf calls log.Logger.Printf()
func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// Println calls log.Logger.Println()
func Println(v ...interface{}) {
	logger.Println(v...)
}
