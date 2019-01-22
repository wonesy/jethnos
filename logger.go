package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger ...
type Logger struct {
	filename string
	*log.Logger
}

var l *Logger
var once sync.Once

// GetLogger ...
func GetLogger() *Logger {
	once.Do(func() {
		l = createLogger()
	})
	return l
}

func (l *Logger) prefixPrint(prefix string, msg string) {
	fmsg := fmt.Sprintf("[%s] %s", prefix, msg)
	l.Output(3, fmsg)
}

// Info ...
func (l *Logger) Info(msg string) {
	l.prefixPrint("INFO", msg)
}

// Warn ...
func (l *Logger) Warn(msg string) {
	l.prefixPrint("WARNING", msg)
}

// Error ...
func (l *Logger) Error(msg string) {
	l.prefixPrint("ERROR", msg)
}

func createLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "NoPants :: ", log.Lshortfile),
	}
}
