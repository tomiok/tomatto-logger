package tomatto

import (
	"encoding/json"
	"errors"
	"fmt"
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
	pretty  bool
	debug   bool
}

func (t *tomatto) marshall(msg interface{}) []byte {
	if t.pretty {
		b, _ := json.MarshalIndent(msg, "", "  ")
		return b
	}
	b, _ := json.Marshal(msg)
	return b
}

//NewTomatto use it at the beginning of the main function
func NewTomatto(pretty, debug bool) {
	// clean tomatto date using setFlag(0)
	log.SetFlags(0)
	infoLogger := log.New(os.Stdout, "", 0)
	warningLogger := log.New(os.Stdout, "", 0)
	errorLogger := log.New(os.Stderr, "", 0)
	_t = &tomatto{
		logInfo: infoLogger,
		logWarn: warningLogger,
		logErr:  errorLogger,
		pretty:  pretty,
		//debug:   debug,
	}
}

func Info(message interface{}) {
	msg := newTomatto(getStackTrace, message, _t.debug)
	b := _t.marshall(msg)
	_t.logInfo.Print(string(b))
}

func Error(message interface{}) {
	msg := newTomatto(getStackTrace, message, _t.debug)
	b := _t.marshall(msg)
	_t.logErr.Print(string(b))
}

func ErrorS(message string, err error) {
	msg := newTomatto(getStackTrace, message+err.Error(), _t.debug)
	b := _t.marshall(msg)
	_t.logErr.Print(string(b))
}

func Warn(message interface{}) {
	msg := newTomatto(getStackTrace, message, _t.debug)
	b := _t.marshall(msg)
	_t.logWarn.Print(string(b))
}

//
// formatted
//

func Infof(s string, values ...interface{}) {
	formatted := fmt.Sprintf(s, values...)
	msg := newTomatto(getStackTrace, formatted, _t.debug)
	b := _t.marshall(msg)
	_t.logInfo.Print(string(b))
}

func Errorf(s string, values ...interface{}) {
	formatted := fmt.Sprintf(s, values...)
	msg := newTomatto(getStackTrace, formatted, _t.debug)
	b := _t.marshall(msg)
	_t.logErr.Print(string(b))
}

func Warnf(s string, values ...interface{}) {
	formatted := fmt.Sprintf(s, values...)
	msg := newTomatto(getStackTrace, formatted, _t.debug)
	b := _t.marshall(msg)
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
	Function string      `json:"function,omitempty"`
	Line     int         `json:"line"`
	Message  interface{} `json:"message,omitempty"`
	// fatal is out of scope
}

func newTomatto(fn func() (uintptr, string, int, error), s interface{}, debug bool) *MsgTomatto {
	pc, file, line, _ := fn()
	var funcCall string
	if debug {
		funcCall = runtime.FuncForPC(pc).Name()
	}

	return &MsgTomatto{
		Line:     line,
		Function: funcCall,
		File:     path.Base(file),
		Message:  s,
	}
}
