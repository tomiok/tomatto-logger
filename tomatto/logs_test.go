package tomatto

import "testing"

func Test_infoLogger(t *testing.T) {
	NewTomatto()
	Info("hello this is a test")
}
