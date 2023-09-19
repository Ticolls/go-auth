package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func newLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)

	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "FROM "+prefix+" -> DEBUG: ", logger.Flags()),
		info:    log.New(writer, "FROM "+prefix+" -> INFO: ", logger.Flags()),
		warning: log.New(writer, "FROM "+prefix+" -> WARNING: ", logger.Flags()),
		err:     log.New(writer, "FROM "+prefix+" -> ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create formate enable logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
