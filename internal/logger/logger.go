package logger

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LevelFatal LogLevel = iota
	LevelError
	LevelInfo
	LevelDebug
)

type Logger struct {
	logger *log.Logger
	level  LogLevel
}

type Logging interface {
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

func New(level LogLevel) *Logger {
	return &Logger{
		logger: log.Default(),
		level:  level,
	}
}

func (l *Logger) Println(args ...interface{}) {
	l.print(args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.print(fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(msg string) {
	if l.level >= LevelDebug {
		l.logger.Println("[DEBUG]: " + msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.level >= LevelInfo {
		l.logger.Println("[INFO]: " + msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.level >= LevelError {
		l.logger.Println("[ERROR]: " + msg)
	}
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *Logger) Level() LogLevel {
	return l.level
}

func (l *Logger) print(args ...interface{}) {
	var prefix string
	switch l.level {
	case LevelFatal:
		prefix = "FATAL"
	case LevelError:
		prefix = "ERROR"
	case LevelInfo:
		prefix = "INFO"
	case LevelDebug:
		prefix = "DEBUG"
	default:
		prefix = "FATAL"
	}

	message := fmt.Sprintf("[%s]: ", prefix) + fmt.Sprint(args...)
	l.logger.Println(message)
}
