package main

import "fmt"

/* Struct is a collection of fields */
type Vertex struct {
	// X int
	// Y int
	X, Y int
}

/* 
 * Struct literal denotes a newly allocated struct
 * value by listing the values of its fields.
 */
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	pp = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}; fmt.Println(v.Y)
	v.Y = 3; fmt.Println(v.Y)

	p := &v
	p.X = 1e9 // p.X == (*p).X when accessing struct fields through pointer
	fmt.Println(v)

	fmt.Println(v1, v2, v3, pp)
}