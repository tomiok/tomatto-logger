package tomatto

import "testing"

func Test_infoLogger(t *testing.T) {
	NewTomatto()
	Info("if you can read this, is working")
}

func Test_warnLogger(t *testing.T) {
	NewTomatto()
	Warn("if you can read this, is working")
}

func Test_errLogger(t *testing.T) {
	NewTomatto()
	Err("if you can read this, is working")
}

func Test_formattedInfo(t *testing.T) {
	NewTomatto()
	Infof("if you can read %s, is %s. And this is a number: %d", "this", "working", 150)
}
