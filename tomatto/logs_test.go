package tomatto

import "testing"

func Test_infoLogger(t *testing.T ) {
	logger := NewTomatto()
	logger.Info("hola esto es una prueba del streaming")

}