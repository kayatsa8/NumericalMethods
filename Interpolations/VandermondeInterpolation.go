package main

import (
	"log"
	"math"
)

func VandermondeInterpolation(points map[float64]float64) (Polynomial, error) {
	A := makeMatrix(points)

	ASolutionResponse := fetch[[]float64]("http://localhost:3004/api/solve_equations/gaussian", A)

	if ASolutionResponse.Err != nil {
		log.Fatalln(ASolutionResponse.Err)

		return EmptyPolynomial(), ASolutionResponse.Err
	}
	
	coefficients := map[int]float64{}

	for degree, coefficient := range ASolutionResponse.Result {
		coefficients[degree] = coefficient
	}

	polynomial := NewPolynomial(coefficients)

	return polynomial, nil
}

func makeMatrix(points map[float64]float64) [][]float64{
	n := len(points)

	A := make([][]float64, n)

	for i := range A {
		A[i] = make([]float64, n + 1)
	}

	i := 0

	for x, y := range points {
		for j := range n {
			A[i][j] = math.Pow(x, float64(j))
		}

		A[i][n] = y

		i++
	}

	return A
}
