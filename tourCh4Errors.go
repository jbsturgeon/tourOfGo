package main

import (
	"fmt"
	"math"
)

const TOL float64 = 1.3e-14
const MAX_ITER int64 = 10000

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	if *e < 0 {
		return fmt.Sprintf("%s %f", "Cannot sqrt negative number: ", *e)
	}

	return "Unkonwn error"
}

func Sqrt(xVal ErrNegativeSqrt) (float64, int64, error) {
	if xVal < 0 {
		return -1, -1, &xVal
	}

	var x float64 = float64(xVal)
	var approx float64 = 0
	var hold float64 = 2 * TOL
	var ii int64 = 0

	for math.Abs(hold-approx) > TOL {
		hold = approx
		approx = approx - ((math.Pow(approx, 2) - x) / (2 * x))
		ii++

		if ii > MAX_ITER {
			break
		}
	}

	return approx, ii, nil
}

func main() {
	vals := []float64{88, -33, 444}

	for _, a := range vals {
		res, iter, err := Sqrt(ErrNegativeSqrt(a))
		if err == nil {
			fmt.Println("Approx Square root of ", a, " is: ", res, "(iter: ", iter, " )")
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
