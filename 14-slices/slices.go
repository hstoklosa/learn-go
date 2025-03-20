package main

import (
	"fmt"
	"strings"
)

/*
 * The type []T is a slice with elements of type T.
 *
 * A slice is a dynamically-sized, flexible view into the elements of an array
 * formed by specifying 2 indices (bounds) a[low : high].
 *
 * A slice does not store any data, it just describes a section of an underlying array.
 *
 * It has both len(s) and cap(s).
 *
 * REF: https://go.dev/blog/slices-intro
 */

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/* Slice literal is like an array without the length */
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, false, true, false, true}
	fmt.Println(r)

	p := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
		{6, true},
	}
	fmt.Println(p)

	/* Slice defaults (0 for low, slice length for high) */
	x := []int{2, 3, 5, 7, 11, 13}
	x = x[1:4]
	fmt.Println(x)

	x = x[:2]
	fmt.Println(x)
	
	x = x[1:]
	fmt.Println(x)

	/* Slice length and capacity */
	y := []int{2, 3, 5, 7, 11, 13}
	printSlice(y)

	y = y[:0] // zero length
	printSlice(y)

	y = y[:4] // extend length
	printSlice(y)

	y = y[2:] // drop first 2 values
	printSlice(y)
	
	/* A nil slice has a length and capacity of 0 and has no underlying array. */
	var z []int
	printSlice(z)
	if z == nil {
		fmt.Println("nil!!!")
	}

	/* 
	 * Creating slice with `make` for a dynamically-sized array. 
	 * It allocates a zeroed array and returns a slice that refers to that array. 
	 */
	c := make([]int, 5) // len(c)=5
	printSlice(c)

	d := make([]int, 0, 5) // len(d)=0, cap(d)=5
	printSlice(d)

	e := d[:2] // len(f)=2, cap(f)=5
	printSlice(e)

	f := e[2:5] // len(f)=3, cap(f)=3
	printSlice(f)

	/* Slices of slices */
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	/* Appending to a slice
	 * If the backing array of s is too small to fit all the given values a bigger 
	 * array will be allocated. The returned slice will point to the newly allocated array.
	 */
	var g []int
	printSlice(g)

	g = append(g, 0)
	printSlice(g)

	g = append(g, 1)
	printSlice(g)

	g = append(g, 2, 3, 4)
	printSlice(g)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}