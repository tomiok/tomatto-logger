package tomatto

import (
	"errors"
	"testing"
)

func Test_infoLogger(t *testing.T) {
	NewTomatto(false, false)
	Info("if you can read this, is working")
}

func Test_warnLogger(t *testing.T) {
	NewTomatto(true, true)
	Warn("if you can read this, is working")
}

func Test_errLogger(t *testing.T) {
	NewTomatto(false, true)
	Error("if you can read this, is working")
}

func Test_formattedInfo(t *testing.T) {
	NewTomatto(true, false)
	Infof("if you can read %s, is %s. And this is a number: %d", "this", "working", 150)
}

func Test_formattedErr(t *testing.T) {
	NewTomatto(true, true)
	ErrorS("if you can read ", errors.New("is working"))
}
