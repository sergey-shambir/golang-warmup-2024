package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, fmt.Errorf("cannot get square root of negative number %v", x)
	}

	z := 1.0
	epsilon := 0.1e-8 * max(x, 1.0)

	for {
		square_miss := z*z - x
		correction := square_miss / (2 * z)
		z -= correction

		if math.Abs(correction) < epsilon {
			return z, nil
		}
	}
}
