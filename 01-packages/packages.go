package main

/*
 * import "fmt"
 * import "math"
 */

import (
	"fmt"
	"math"

	/*
	 * By convention, the package name is the same as the last element of the import path.
	 * For instance, the "math/rand" package comprises files that begin with the statement package rand.
	 */
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
	fmt.Printf("Now you have %b problems. \n", math.Sqrt(7))

	/* A name is exported if it begins with a capital letter. */
	fmt.Println(math.Pi)

	/*
	 * Any "unexported" names are not accessible from outside the package.
	 * fmt.Println(math.pi)
	 */
}