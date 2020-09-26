package tomatto

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	iLog *log.Logger
)

type tomatto struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewTomatto() {
	infoLogger := log.New(os.Stdout, "", 0)
	warningLogger := log.New(os.Stdout, "", 0)
	errorLogger := log.New(os.Stdout, "", 0)
	t := &tomatto{
		infoLogger:    infoLogger,
		warningLogger: warningLogger,
		errorLogger:   errorLogger,
	}
	iLog = t.infoLogger
}

func Info(message interface{}) {
	pc, file, line, _ := getStackTrace()

	msg := &MsgTomatto{
		Line:    line,
		Method:  runtime.FuncForPC(pc).Name(),
		File:    path.Base(file),
		Message: message,
	}

	b, _ := json.MarshalIndent(msg, "", "    ")
	iLog.Print(string(b))
}

func getStackTrace() (uintptr, string, int, error) {
	pc, file, line, ok := runtime.Caller(2)

	if !ok {
		return 0, "", 0, errors.New("cannot get stack trace")
	}

	return pc, file, line, nil
}

type MsgTomatto struct {
	Line    int         `json:"line"`
	Method  string      `json:"method"`
	File    string      `json:"file"`
	Message interface{} `json:"message,omitempty"`
	Err     interface{} `json:"error,omitempty"`
	Warn    interface{} `json:"warn,omitempty"`
	// fatal is out of scope
}
