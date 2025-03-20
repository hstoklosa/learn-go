package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	/* Init and post statements are optional */
	for ;sum < 990; {
		sum += 1
	}
	fmt.Println(sum)

	/* While loops */
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	/* Infinite loops by omitting loop condition */
	for { fmt.Println("Hello World") }
}