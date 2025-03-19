package main

import (
	"fmt"
	"math"
)

func main3() {
	/* A name is exported if it begins with a capital letter. */
	fmt.Println(math.Pi)
	
	/*
	 * Any "unexported" names are not accessible from outside the package.
	 * fmt.Println(math.pi)
	 */
}