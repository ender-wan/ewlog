package ewlog_test

import (
	"fmt"
	"os"

	"github.com/ender-wan/ewlog"
)

// Init log and output log
func ExampleAddLogOutput() {
	logfile, err := os.OpenFile("LogFile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logfile.Close()

	ewlog.SetLogLevel(1)
	ewlog.AddLogOutput(logfile)

	ewlog.Info("ewlog")
}
