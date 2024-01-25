package logger

import (
	"log"
	"os"
)

func NewLogger() (infoLogger *log.Logger, errLogger *log.Logger, err error) {
	logfile, err := os.OpenFile("auth.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, err
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile

	log.SetOutput(logfile)
	log.SetFlags(flags)

	infoLogger = log.New(logfile, "INFO: ", flags)
	errLogger = log.New(logfile, "ERROR: ", flags)

	return
}
