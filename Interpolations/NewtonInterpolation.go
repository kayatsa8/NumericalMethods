package main

import (
	"fmt"
	"log"
)

// gets a map of <x, y> points
func newtonInterpolation(points map[float64]float64) (Polynomial, error) {
	if len(points) == 0 {
		return EmptyPolynomial(), fmt.Errorf("the points collection is empty")
	}

	p, xPrev := makeP0(points)
	denominatorPolynomial := NewPolynomial(map[int]float64{
		0: 1,
	})

	for x, y := range points {
		numerator := y - calculatePolynomailValue(&p, x)

		denominatorPolynomial = makeDenominator(&denominatorPolynomial, xPrev)
		denominator := calculatePolynomailValue(&denominatorPolynomial, x)

		c := numerator / denominator

		additionPolynomial := multipliePolynomials([]Polynomial{denominatorPolynomial}, c)

		p = addPolynomials(&p, &additionPolynomial)
			
		xPrev = x
	}

	return p, nil
}

func makeP0(points map[float64]float64) (Polynomial, float64) {
	p0 := EmptyPolynomial()
	_x := 0.0

	for x, y := range points {
		p0 = NewPolynomial(map[int]float64{
			0: y,
		})

		_x = x

		break
	}

	delete(points, _x)

	return p0, _x
}

func calculatePolynomailValue(polynomial *Polynomial, value float64) float64 {
	response := fetch("http://localhost:3000/api/polynomials/calculate")

	if response.Err != nil {
		log.Fatal(response.Err)
		return 0
	}

	return response.Result.(float64)
}

func makeDenominator(denominator *Polynomial, xPrev float64) Polynomial {
	mult := NewPolynomial(map[int]float64{
		1: 1,
		0: xPrev,
	})

	return multipliePolynomials([]Polynomial{mult, *denominator}, 1)
}

func multipliePolynomials(polynomials []Polynomial, coefficients float64) Polynomial {
	response := fetch("http://localhost:3000/api/polynomials/multiplie")

	if response.Err != nil {
		log.Fatal(response.Err)
		return EmptyPolynomial()
	}

	return response.Result.(Polynomial)
}

func addPolynomials(pol1 *Polynomial, pol2 *Polynomial) Polynomial {
	response := fetch("http://localhost:3000/api/polynomials/add")

	if response.Err != nil {
		log.Fatal(response.Err)
		return EmptyPolynomial()
	}

	return response.Result.(Polynomial)
}


