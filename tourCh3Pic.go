//Implement Pic. It should return a slice of length dy, each element of which
// is a slice of dx 8-bit unsigned integers. When you run the program, it will
// display your picture, interpreting the integers as grayscale (well, bluescale)
// values.
//
//The choice of image is up to you. Interesting functions include:
// (x+y)/2, x*y, and x^y.
//
//You need to use a loop to allocate each []uint8 inside the [][]uint8.
//
//Use uint8(intValue) to convert between types.
package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	values := make([][]uint8, dy)

	for ii := 0; ii < dy; ii++ {
		values[ii] = make([]uint8, dx)

		for jj := 0; jj < dx; jj++ {
			values[ii][jj] = uint8(ii * jj)
		}
	}

	return values
}

func main() {
	pic.Show(Pic)
}
