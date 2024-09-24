package main

import (
	"fmt"
)


// gets a map of <x, y> points
func LagrangeInterpolation(points map[float64]float64) (Polynomial, error) {
	if len(points) == 0 {
		return EmptyPolynomial(), fmt.Errorf("the points collection is empty")
	}

	P := EmptyPolynomial()
	diffPolinomials := makeDiffPolynomials(points)

	for x, y := range points {
		L_i := makeL_i(points, x, diffPolinomials)
		mul := multipliePolynomials([]Polynomial{L_i}, y)

		P = addPolynomials(&P, &mul)
	}

	return P, nil
}

func makeDiffPolynomials(points map[float64]float64) map[float64]Polynomial {
	var diffPolynomials map[float64]Polynomial = map[float64]Polynomial{}

	for x := range points {
		diffPolynomials[x] = NewPolynomial(map[int]float64{
			0: -x,
			1: 1,
		})
	}

	return diffPolynomials
}


func makeL_i(points map[float64]float64, x_i float64, diffPolynomials map[float64]Polynomial) Polynomial {
	C_i := makeC_i(points, x_i)

	l := NewPolynomial(map[int]float64{
		0: 1,
	})

	for x, pol := range diffPolynomials{
		if x != x_i {
			l = multipliePolynomials([]Polynomial{l, pol}, 1)
		}
	}

	l = multipliePolynomials([]Polynomial{l}, C_i)

	return l
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


