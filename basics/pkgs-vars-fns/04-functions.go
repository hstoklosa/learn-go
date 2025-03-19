package main

import "fmt"

// https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main4() {
	fmt.Println(add(10, 15))
	fmt.Println(multiply(2, 4))
}