package logger

import (
	"fmt"
	"log"
)

const (
	LevelDebug int = iota
	LevelInfo
	LevelError
	LevelFatal
)

type Logger struct {
	logger *log.Logger
	level  int
}

type Logging interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

func New(level int) *Logger {
	return &Logger{
		logger: log.Default(),
		level:  level,
	}
}

func (l *Logger) Log(args ...interface{}) {
	l.print(args...)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	l.printf(format, args...)
}

func (l *Logger) Debug(msg string) {
	if l.level >= LevelDebug {
		l.logger.SetPrefix("[DEBUG]: ")
		l.logger.Println(msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.level >= LevelInfo {
		l.logger.SetPrefix("[INFO]: ")
		l.logger.Println(msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.level >= LevelError {
		l.logger.SetPrefix("[ERROR]: ")
		l.logger.Println(msg)
	}
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *Logger) Level() int {
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

	l.logger.SetPrefix(fmt.Sprintf("[%s]: ", prefix))
	l.logger.Println(args...)
}

func (l *Logger) printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	l.print(s)
}
