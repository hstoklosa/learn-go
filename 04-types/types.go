package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
 * bool
 * string
 *
 * int  int8  int16  int32  int64
 * uint uint8 uint16 uint32 uint64 uintptr
 *
 * byte // alias for uint8
 * rune // alias for int32, represents a Unicode code point
 *
 * float32 float64
 * complex64 complex128
 */

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/* Zero values */
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println(i, f, b, s)

	/* T(v) converts the value v to the type T. */
	var x, y int = 3, 5
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var u uint = uint(f)

	fmt.Println(x, y, f2, u)

	/* Type Inference */
	v := 45
	fmt.Printf("v is of type %T", v)
}