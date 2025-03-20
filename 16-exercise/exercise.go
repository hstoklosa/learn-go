package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
		
		for j := 0; j < dx; j++ {
			fn := (i + j)/2
			p[i][j] = uint8(fn)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}