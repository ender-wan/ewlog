// Package ewlog provides a simple log tool for multiple output
package ewlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
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
		logs:     []*log.Logger{log.New(os.Stdout, "", log.LstdFlags)},
		logLevel: INFO,
	}
}

func logHead(level int) string {
	var lv string
	switch level {
	case DEBUG:
		lv = "[Debug]"
	case INFO:
		lv = "[Info]"
	case WARN:
		lv = "[Warn]"
	case ERROR:
		lv = "[Error]"
	case FATAL:
		lv = "[Fatal]"
	default:
		lv = "[Unknown]"
	}

	funcname := ""

	function, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	} else {
		funcname = runtime.FuncForPC(function).Name()
	}

	file = shortFile(file)

	head := fmt.Sprintf("%s:%d %s %s - ", file, line, funcname, lv)

	return head
}

func shortFile(file string) string {
	for i := len(file) - 1; i >= 0; i-- {
		if file[i] == '/' {
			return file[i+1:]
		}
	}

	return ""
}

// SetLogLevel set output log level
func SetLogLevel(level int) {
	ewlog.logLevel = level
}

// AddLogOutput add a writer to output log
func AddLogOutput(out io.Writer) {
	elog := log.New(out, "", log.LstdFlags)
	elog.SetOutput(out)
	ewlog.logs = append(ewlog.logs, elog)
}

func output(level int, v ...interface{}) {
	if ewlog.logLevel > level {
		return
	}

	head := logHead(level)
	v = append([]interface{}{head}, v...)

	for _, tlog := range ewlog.logs {
		tlog.Print(v...)
	}
}

func outputf(level int, format string, v ...interface{}) {
	if ewlog.logLevel > level {
		return
	}

	head := logHead(level)

	for _, tlog := range ewlog.logs {
		tlog.Printf(head+format, v...)
	}
}

func Debug(v ...interface{}) {
	output(DEBUG, v...)
}

func Debugf(format string, v ...interface{}) {
	outputf(DEBUG, format, v...)
}

func Info(v ...interface{}) {
	output(INFO, v...)
}

func Infof(format string, v ...interface{}) {
	outputf(INFO, format, v...)
}

func Warn(v ...interface{}) {
	output(WARN, v...)
}

func Warnf(format string, v ...interface{}) {
	outputf(WARN, format, v...)
}

func Error(v ...interface{}) {
	output(ERROR, v...)
}

func Errorf(format string, v ...interface{}) {
	outputf(ERROR, format, v...)
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	output(FATAL, v...)
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	outputf(FATAL, format, v...)
	os.Exit(1)
}
