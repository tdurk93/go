package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var z [][]uint8
	z = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		z[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			z[i][j] = uint8(i * j)
		}
	}
	return z
}

func main() {
	pic.Show(Pic)
}

