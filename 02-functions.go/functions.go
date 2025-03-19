package main

import "fmt"

// https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

/*
 * A return statement without arguments returns the named return values. This is known
 * as a "naked" return. Naked return statements should be used only in short functions,
 */
 func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(10, 15))
	fmt.Println(multiply(2, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d := split(17)
	fmt.Println(c, d)
}