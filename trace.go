package sdk

import (
	"io/ioutil"
	"log"
)

// Loggers
var (
	ERROR    *log.Logger
	CRITICAL *log.Logger
	WARN     *log.Logger
	INFO     *log.Logger
	DEBUG    *log.Logger
)

func initLogging() {
	ERROR = log.New(ioutil.Discard, "", 0)
	CRITICAL = log.New(ioutil.Discard, "", 0)
	WARN = log.New(ioutil.Discard, "", 0)
	INFO = log.New(ioutil.Discard, "", 0)
	DEBUG = log.New(ioutil.Discard, "", 0)
}
