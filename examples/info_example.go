package main

import (
	"fmt"
	"github.com/tomiok/tomatto-logger/tomatto"
)

func main() {
	tomatto.NewTomatto(false, false)

	sum(165, 983)

}

func sum(a, b int) int {
	tomatto.Info(fmt.Sprintf("adding %d and %d", a, b))

	res := a + b

	tomatto.Info(fmt.Sprintf("the result is %d", res))
	return res
}
