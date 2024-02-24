package logger

import (
	"io"
	"os"

	"github.com/google/logger"
)

var errLog *logger.Logger

// N.B. takes over stderr so the google logger doesn't print to user
var stdErr *os.File

func init() {
	errLog = logger.Init("thetundra", false, true, io.Discard)
	stdErr = os.Stderr
	os.Stderr, _ = os.Open(os.DevNull)
	//errLog.
}

func Log(err error) {
	errLog.Warningln(err)
}

// restores stderr and closes the logger
// Note that the logger sends stderr to /dev/null
// by default
func Close() {
	os.Stderr = stdErr
	errLog.Close()
}
