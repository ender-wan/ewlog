// Package ewlog provides a simple log tool for multiple output
package ewlog

import (
	"io"
	"log"
	"os"
)

// These flags define which log will be printed
const (
	// DEBUG level
	DEBUG = 1 << iota
	INFO
	WARN
	ERROR
	FATAL
)

type ewlogT struct {
	logs     []*log.Logger
	logLevel int
}

var ewlog *ewlogT

func init() {
	ewlog = &ewlogT{
		logs:     []*log.Logger{log.New(os.Stdout, "", 3)},
		logLevel: INFO,
	}
}

// SetLogLevel set output log level
func SetLogLevel(level int) {
	ewlog.logLevel = level
}

// AddLogOutput add a writer to output log
func AddLogOutput(out io.Writer) {
	elog := log.New(out, "", 3)
	elog.SetOutput(out)
	ewlog.logs = append(ewlog.logs, elog)
}

func Debug(v interface{}) {
	if ewlog.logLevel > DEBUG {
		return
	}
	for _, tlog := range ewlog.logs {
		tlog.Printf("Debug %v", v)
	}
}

func Debugf(format string, v ...interface{}) {
	if ewlog.logLevel > DEBUG {
		return
	}
	for _, tlog := range ewlog.logs {
		tlog.Printf("Debug "+format, v...)
	}
}

func Info(v interface{}) {
	if ewlog.logLevel > INFO {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Printf("Info %v", v)
	}
}

func Infof(format string, v ...interface{}) {
	if ewlog.logLevel > INFO {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Printf("Info "+format, v...)
	}
}

func Warn(v interface{}) {
	if ewlog.logLevel > WARN {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Printf("Warn %v", v)
	}
}

func Warnf(format string, v ...interface{}) {
	if ewlog.logLevel > WARN {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Printf("Warn "+format, v...)
	}
}

func Error(v interface{}) {
	if ewlog.logLevel > ERROR {
		return
	}
	for _, tlog := range ewlog.logs {
		tlog.Printf("Error %v", v)
	}
}

func Errorf(format string, v ...interface{}) {
	if ewlog.logLevel > ERROR {
		return
	}
	for _, tlog := range ewlog.logs {
		tlog.Printf("Error "+format, v...)
	}
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func Fatal(v interface{}) {
	if ewlog.logLevel > FATAL {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Fatalf("Fatal %v", v)
	}
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	if ewlog.logLevel > FATAL {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Fatalf("Fatal "+format, v...)
	}
}
