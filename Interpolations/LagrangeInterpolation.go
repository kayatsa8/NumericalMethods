package main

import (
	"fmt"
	"log"
)


// gets a map of <x, y> points
func LagrangeInterpolation(points map[float64]float64) (Polynomial, error) {
	return EmptyPolynomial(), nil
}




func makeC_i(points map[float64]float64, x_i float64) float64 {
	var mul float64 = 1

	for x := range points {
		if x != x_i {
			mul *= (x_i - x)
		}
	}

	return 1 / mul
}


