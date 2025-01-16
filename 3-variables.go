package main

import "fmt"

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d, e = 1, 2
	fmt.Println(d, e)

	var f = true
	fmt.Println(f)

	var g int
	fmt.Println(g)

	h := "apple" // var h string = "apple"
	fmt.Println(h)
}