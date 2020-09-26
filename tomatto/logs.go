package tomatto

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

type Tomatto struct {
	infoLogger *log.Logger
}

func NewTomatto() *Tomatto {
	infoLogger := log.New(os.Stdout, "", 0)
	return &Tomatto{
		infoLogger: infoLogger,
	}
}

func (l *Tomatto) Info(message interface{}) {
	pc, file, line, _ := getStackTrace()

	msg := &MsgTomatto{
		Line:    line,
		Method:  runtime.FuncForPC(pc).Name(),
		File:    path.Base(file),
		Message: message,
	}

	b, _ := json.MarshalIndent(msg, "", "    ")
	l.infoLogger.Print(string(b))
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
