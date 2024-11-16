package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	debugLog *log.Logger
}

func New() *Logger {
	flags := log.LstdFlags | log.Lmicroseconds | log.Lshortfile | log.LUTC

	return &Logger{
		infoLog:  log.New(os.Stdout, "[INFO] ", flags),
		errorLog: log.New(os.Stderr, "[ERROR] ", flags),
		debugLog: log.New(os.Stdout, "[DEBUG] ", flags),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLog.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLog.Printf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.debugLog.Printf(format, v...)
}
