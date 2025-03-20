package main

import "fmt"

/*
 * var p *int (type *T is a pointer to a T value - zero value is nil)
 * i := 42; p = &i (generates a pointer to its operand)
 * fmt.Println(*p); *p = 21 (denotes the pointer's underlying value)
 */

func main() {
	i, j := 42, 2701

	p := &i			// point to i
	fmt.Println(*p) // read i through pointer
	*p = 21			// set i through pointer
	fmt.Println(i)  // see value of i

	p = &j			// point to j
	*p = *p / 37 	// divide j through pointer
	fmt.Println(j)	// see value of j
}