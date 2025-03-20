package main

import "fmt"

/*
 * The type [n]T is an array of n values of type T.
 *
 * The expression `var a [10]int` delcares a variable as an array of 10 integers.
 * An array's length is part of its type, so arrays can't be resized.
 */

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Array"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}