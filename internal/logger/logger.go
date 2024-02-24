package logger

import "github.com/google/logger"

var errLog *logger.Logger

func init() {
	errLog = logger.Init("InternalErrors", false, true, nil)
}

func Log(err error) {
	errLog.Errorln(err)
}
