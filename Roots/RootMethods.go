package main

import (
	"fmt"
	"math"
)

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

func regulaFalsi(a float64, b float64, f func(float64) float64, delta float64) (float64, error) {
	// same sign
	if f(a) * f(b) > 0 {
		return 0, fmt.Errorf("both bounds have the same sign")
	}

	start := a
	end := b
	mid := 0.0
	m := 0.0
	fStart := 0.0

	for end - start > delta {
		fStart = f(start)

		m = (f(end) - fStart) / (end - start)
		mid = start - (fStart / m)

		if f(mid) == 0 {
			return mid, nil
		} else if f(mid) * fStart > 0 {
			start = mid
		} else {
			end = mid
		}
	}

	return mid, nil
}

func fixedPoint(guess float64, g func(float64) float64, delta float64) float64 {
	curr := guess
	next := g(curr)

	for math.Abs(next-curr) > delta {
		curr = next
		next = g(curr)
	}

	return next
}

func newtonMethod(guess float64, f func(float64) float64, fTag func(float64) float64, delta float64) float64 {
	curr := guess
	next := curr - f(curr) / fTag(curr)

	for math.Abs(next - curr) > delta {
		curr = next
		next = curr - f(curr) / fTag(curr)
	}

	return next
}