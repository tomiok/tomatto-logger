package main

import (
	"fmt"
	"github.com/tomiok/tomatto-logger/tomatto"
)

func main() {
	tomatto.NewTomatto()

	sum(165, 983)

	tomatto.Test("hello %s %d", "tomas", 1)

}

func sum(a, b int) int {
	tomatto.Info(fmt.Sprintf("adding %d and %d", a, b))

	res := a + b

	tomatto.Info(fmt.Sprintf("the result is %d", res))
	return res
}
