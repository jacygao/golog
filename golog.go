/*
Package golog is a very basic logging library that wraps the standard Go log package.
golog implements logr.Logger interface and provides basic leveled logging and structured tags.
golog can be used as an alternative logging library when the Go log package is a dependency.*/
package golog

import (
	"fmt"
	"log"
	"os"
)

// Logr Levels
const (
	DEBUG = iota
	TRACE
	WARN
	ERROR
	FATAL
)

// Logger defines the instance of golog
type Logger struct {
	// The level at which this logger will log.  Can only be set on New
	level int
	tag   []Tag
}

// Tag stores tag strings in key-value format
type Tag struct {
	key   string
	value interface{}
}

// New initialises a new Logger instance
func New(lvl int) *Logger {
	return &Logger{
		level: lvl,
	}
}

func (l *Logger) message(out string, lv int) {
	if l.level <= lv {
		if len(l.tag) > 0 {
			log.Print(out + " " + fmt.Sprintln(l.tag))
		} else {
			log.Print(out)
		}
	}
}

// Debug prints log messages on Debug or lower log level
func (l *Logger) Debug(msg ...interface{}) {
	l.message("INFO:  "+fmt.Sprint(msg...), DEBUG)
}

// Debugf prints formatted log messages on Debug or lower log level
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.message("INFO:  "+fmt.Sprintf(format, args...), DEBUG)
}

// Debugw prints log message with key value pair on Debug or lower log level
func (l *Logger) Debugw(msg string, kv ...interface{}) {
	l.With(kv...).Debug(msg)
}

// Info prints log messages on Info or lower log level
func (l *Logger) Info(info ...interface{}) {
	l.message("INFO:  "+fmt.Sprint(info...), DEBUG)
}

// Infof prints formatted log messages on Info or lower log level
func (l *Logger) Infof(format string, info ...interface{}) {
	l.message("INFO:  "+fmt.Sprintf(format, info...), DEBUG)
}

// Infow prints log message with key value pair on Info or lower log level
func (l *Logger) Infow(msg string, kv ...interface{}) {
	l.With(kv...).Info(msg)
}

// Warn prints log messages on Warn or lower log level
func (l *Logger) Warn(warn ...interface{}) {
	l.message("WARN:  "+fmt.Sprint(warn...), WARN)
}

// Warnf prints formatted log messages on Warn or lower log level
func (l *Logger) Warnf(format string, warn ...interface{}) {
	l.message("WARN:  "+fmt.Sprintf(format, warn...), WARN)
}

// Warnw prints log message with key value pair on Warn or lower log level
func (l *Logger) Warnw(msg string, kv ...interface{}) {
	l.With(kv...).Warn(msg)
}

// Warn prints log messages on Error or lower log level
func (l *Logger) Error(err ...interface{}) {
	l.message("ERROR: "+fmt.Sprint(err...), ERROR)
}

// Errorf prints formatted log messages on Error or lower log level
func (l *Logger) Errorf(format string, err ...interface{}) {
	l.message("ERROR: "+fmt.Sprintf(format, err...), ERROR)
}

// Errorw prints log message with key value pair on Error or lower log level
func (l *Logger) Errorw(msg string, kv ...interface{}) {
	l.With(kv...).Error(msg)
}

// Fatal prints log messages and then calls os.Exit
func (l *Logger) Fatal(msg ...interface{}) {
	l.message(fmt.Sprint(msg...), FATAL)
	os.Exit(1)
}

// Fatalf prints formatted log messages and then calls os.Exit
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.message(fmt.Sprintf(format, args...), FATAL)
	os.Exit(1)
}

// Fatalw prints log message with key value pair and then calls os.Exit
func (l *Logger) Fatalw(msg string, kv ...interface{}) {
	l.With(kv...).Fatal(msg)
}

// With gets a child logger instance with specific key value fields attached.
func (l *Logger) With(kvs ...interface{}) *Logger {
	childLogger := l.clone()
	return childLogger.with(kvs...)
}

func (l *Logger) with(args ...interface{}) *Logger {
	for i := 0; i < len(args); {
		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			l.Warn("ignored key value pairs in tags: %v", args[i])
			break
		}
		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			l.Warn("invalid key type in tags: %v", key)
		} else {
			l.tag = append(l.tag, Tag{keyStr, val})
		}
		i += 2
	}
	return l
}

func (l *Logger) clone() *Logger {
	copy := *l
	return &copy
}
