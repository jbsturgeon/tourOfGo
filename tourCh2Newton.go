package main

import (
	"fmt"
	"math"
)

const TOL float64 = 1.e-14
const MAX_ITER int64 = 1000

func Sqrt(x float64) (float64, float64, int64) {
	exact := math.Sqrt(x)

	var approx float64 = 0
	var hold float64 = 2 * TOL
	var ii int64 = 0

	//	for ii := 0; ii < 100; ii++ {
	for math.Abs(hold-approx) > TOL {
		hold = approx
		approx = approx - ((math.Pow(approx, 2) - x) / (2 * x))
		ii++

		if ii > MAX_ITER {
			break
		}
	}

	return approx, exact, ii
}

func main() {
	testNum := float64(33)

	var app, ex float64 = 0, 0
	var iter int64 = 0

	app, ex, iter = Sqrt(testNum)

	fmt.Println("Approx Square root of ", testNum, " is: ", app, "(iter: ", iter, " )")
	fmt.Println("Exact value is: ", ex)
}
