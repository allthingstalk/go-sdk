package sdk

import (
	"io/ioutil"
	"log"
)

// Loggers
var (
	ERROR *log.Logger
	INFO  *log.Logger
	DEBUG *log.Logger
)

func init() {
	ERROR = log.New(ioutil.Discard, "", 0)
	INFO = log.New(ioutil.Discard, "", 0)
	DEBUG = log.New(ioutil.Discard, "", 0)
}
