package ewlog

import (
	"io"
	"log"
	"os"
)

const (
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

func InitLog(out io.Writer, level int) {
	elog := log.New(out, "", 3)
	elog.SetOutput(out)
	ewlog.logs = append(ewlog.logs, elog)
	ewlog.logLevel = level
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

func Fatal(v interface{}) {
	if ewlog.logLevel > FATAL {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Fatalf("Fatal %v", v)
	}
}

func Fatalf(format string, v ...interface{}) {
	if ewlog.logLevel > FATAL {
		return
	}

	for _, tlog := range ewlog.logs {
		tlog.Fatalf("Fatal "+format, v...)
	}
}
