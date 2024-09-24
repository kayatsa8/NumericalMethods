package main

import (
	"fmt"
)


// points = <x, f(x)>; diffPoints = <x, f'(x)>
func HermiteInterpolation(points map[float64]float64, diffPoints map[float64]float64) (Polynomial, error){
	if len(points) == 0 || len(diffPoints) == 0 || len(points) != len(diffPoints) {
		return EmptyPolynomial(), fmt.Errorf("bad points/diffPoints collection")
	}

	sumH := EmptyPolynomial()
	sumHd := EmptyPolynomial()

	diffPolinomials := makeDiffPolynomials(points)

	for x := range points {
		L := makeL_i(points, x, diffPolinomials)
		LTag := derivative(&L)
		L = multipliePolynomials([]Polynomial{L, L}, 1)
		
		Hc := NewPolynomial(map[int]float64{
			0: 2 * x,
			1: -2,
		})
		Hc = multipliePolynomials([]Polynomial{Hc, LTag}, 1)
		one := NewPolynomial(map[int]float64{0: 1})
		Hc = addPolynomials(&Hc, &one)

		Hdc := NewPolynomial(map[int]float64{
			0: -x,
			1: 1,
		})

		H := multipliePolynomials([]Polynomial{Hc, L}, points[x])
		Hd := multipliePolynomials([]Polynomial{Hdc, L}, diffPoints[x])

		sumH = addPolynomials(&sumH, &H)
		sumHd = addPolynomials(&sumHd, &Hd)
	}

	P := addPolynomials(&sumH, &sumHd)

	return P, nil
}












