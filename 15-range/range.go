package main

import "fmt"

/*
 * The range form of the for loop iterates over a slice or map.
 * When ranging over a slice, two values are returned for each iteration (i, v).
 */

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
 
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	localPow := make([]int, 10)
	for i := range localPow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, v := range localPow {
		fmt.Printf("%d\n", v)
	}
}