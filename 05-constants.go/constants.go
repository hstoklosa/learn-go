package main

import "fmt"

/*
 * Cannot be declared using the := syntax
 * const Pi = 3.14
 */

 const (
	Pi = 3.14
)

/*
 * Numeric constants are high-precision values, and
 * untyped constant takes the type needed by its context.
 */

 const (
	Big = 1 << 100 // binary 1 followed by 100 zeroes
	Small = 1 >> 99 // 1 << 1, or 2
)

func needInt(x int) int {
	return (x * 10) + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "Swiat"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}