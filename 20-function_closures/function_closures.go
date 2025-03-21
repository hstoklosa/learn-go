package main

import "fmt"

/* Go functions may be closures.
 *
 * A closure is a function value that references variables from outside its body. It may
 * access and assign to the referenced variables; the function is "bound" to the variables.
 */

func adder() func(int) int {
	sum := 0
	return func(x int) int { // the closure
		sum += x
		return sum
	}
}
 
func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}