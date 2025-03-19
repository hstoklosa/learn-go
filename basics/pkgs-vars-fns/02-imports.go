package main

/*
 * import "fmt"
 * import "math"
 *
 * Better style is to use the factored import statement.
 */

import (
	"fmt"
	"math"
)

func main2() {
	fmt.Printf("Now you have %b problems. \n", math.Sqrt(7))
}