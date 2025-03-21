package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/* A nil map has no keys, nor can keys be added. */
var m map[string]Vertex

/* Map literals require keys. */
var _m = map[string]Vertex{
	// "Base": Vertex{40.68433, -74.39967}
	"Base": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Base"] = Vertex{40.68433, -74.39967}

	fmt.Println(m["Base"])
	fmt.Println(_m)

	__m := make(map[string]int)

	__m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	__m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(__m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := __m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}