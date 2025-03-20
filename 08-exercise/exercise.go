package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}

	return z
}

func PreciseSqrt(x float64) float64 {
	z := 1.0
	const tol = 1e-10
	for {
		nextZ := z - (z*z - x) / (2*z)
		if math.Abs(nextZ-z) < tol {
			return nextZ
		}
		z = nextZ
	}
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(PreciseSqrt(4))
}