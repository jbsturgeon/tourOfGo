package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	x := Vertex{3, 4}
	v := x //tye with and without the & to show x is not modified
	var y *Vertex = &Vertex{8, 2}

	v.Scale(3)
	fmt.Println("x - ", x, x.Abs())
	fmt.Println("v - ", v, v.Abs())

	fmt.Println("\n*y - ", *y, y.Abs())
	fmt.Println("\n... scaling y by 4...")
	y.Scale(4)
	fmt.Println("\n*y - ", *y, y.Abs())
}
