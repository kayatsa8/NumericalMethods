package main

import "fmt"

func bisection(a float64, b float64, f func(float64) float64, delta float64) (float64, error) {
	// same sign
	if f(a) * f(b) > 0 {
		return 0, fmt.Errorf("both bounds have the same sign")
	}

	start := a
	end := b
	mid := (start + end) / 2
	var fMid float64

	for end-start > delta {
		mid = (start + end) / 2

		fMid = f(mid)

		if fMid == 0 {
			return mid, nil
		} else if fMid * f(start) > 0 {
			start = mid
		} else {
			end = mid
		}
	}

	return mid, nil
}