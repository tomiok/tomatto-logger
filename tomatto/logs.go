package tomatto

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

type iLog struct {
	logInfo   *log.Logger
	logWarn   *log.Logger
	logErr    *log.Logger
	logString string
}

type tomatto struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewTomatto() *iLog {
	// clean ILog date using setFlag(0)
	log.SetFlags(0)
	infoLogger := log.New(os.Stdout, "", 0)
	warningLogger := log.New(os.Stdout, "", 0)
	errorLogger := log.New(os.Stdout, "", 0)
	t := &tomatto{
		infoLogger:    infoLogger,
		warningLogger: warningLogger,
		errorLogger:   errorLogger,
	}
	return &iLog{
		logInfo: t.infoLogger,
		logWarn: t.warningLogger,
		logErr:  t.errorLogger,
	}
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
	l := &iLog{
		logString: string(b),
	}
	log.Print(l.logString)
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
