package tomatto

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

var _t *tomatto

type tomatto struct {
	logInfo *log.Logger
	logWarn *log.Logger
	logErr  *log.Logger
}

//NewTomatto use it at the beginning of the main function
func NewTomatto() {
	// clean tomatto date using setFlag(0)
	log.SetFlags(0)
	infoLogger := log.New(os.Stdout, "", 0)
	warningLogger := log.New(os.Stdout, "", 0)
	errorLogger := log.New(os.Stderr, "", 0)
	_t = &tomatto{
		logInfo: infoLogger,
		logWarn: warningLogger,
		logErr:  errorLogger,
	}
}

func Info(message interface{}) {
	msg := newTomatto(getStackTrace, message)
	b, _ := json.MarshalIndent(msg, "", "  ")
	_t.logInfo.Print(string(b))
}

func Err(message interface{}) {
	msg := newTomatto(getStackTrace, message)
	b, _ := json.MarshalIndent(msg, "", "  ")
	_t.logWarn.Print(string(b))
}

func Warn(message interface{}) {
	msg := newTomatto(getStackTrace, message)
	b, _ := json.MarshalIndent(msg, "", "  ")
	_t.logWarn.Print(string(b))
}

func getStackTrace() (uintptr, string, int, error) {
	pc, file, line, ok := runtime.Caller(3)

	if !ok {
		return 0, "", 0, errors.New("cannot get stack trace")
	}

	return pc, file, line, nil
}

//MsgTomatto is the base structure for the JSON marshalled, contains all the info
//necessary to build the correct and easy-to-read log.
//File: is the file where the log is called.
//Function: is the name like {package}.{function}
//Message: the actual message to be logged.
type MsgTomatto struct {
	File     string      `json:"file"`
	Function string      `json:"function"`
	Line     int         `json:"line"`
	Message  interface{} `json:"message,omitempty"`
	// fatal is out of scope
}

func newTomatto(fn func() (uintptr, string, int, error), message interface{}) *MsgTomatto {
	pc, file, line, err := fn()

	if err != nil {
		return nil
	}

	return &MsgTomatto{
		Line:     line,
		Function: runtime.FuncForPC(pc).Name(),
		File:     path.Base(file),
		Message:  message,
	}
}
