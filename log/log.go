package log

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
)

// Log for lumberjack
type Log struct {
	*log.Logger
}

var logInfo *Log
var syncLog sync.Once

// Get : get log with lumberjack
func Get() *Log {
	syncLog.Do(func() {
		logInfo = loadLog()
	})
	return logInfo
}

func loadLog() *Log {
	e, err := os.OpenFile("/apps/2pay/billerpayment-service/logs/2pay.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		log.Panic(err.Error())
	}
	logFile := log.New(e, "", log.Ldate|log.Ltime)
	logFile.SetOutput(&lumberjack.Logger{
		Filename:   "/apps/2pay/billerpayment-service/logs/2pay.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
		Compress:   true,
	})

	return &Log{logFile}
}
