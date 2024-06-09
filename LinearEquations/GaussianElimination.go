package main

import (
	"fmt"
	"math"
)

func GaussianElimination(A [][]float64) ([]float64, error) {
	if !validateInput(A) {
		return []float64{}, fmt.Errorf("the input is invalid")
	}

	for i := range A {
		scale(A, i)

		pivot(A, i)

		subtruct(A, i)
	}

	X := solve(A)

	return X, nil
}


func validateInput(A [][]float64) bool {
	if A == nil {
		return false
	}

	n := len(A)
	m := len(A[0])

	if n + 1 != m {
		return false
	}

	for _, line := range A {
		if len(line) != m {
			return false
		}
	}

	return true
}

func scale(A [][]float64, i int) {
	for k := i; k < len(A); k++ {
		_max := getMax(A, k)

		multiplieLine(A[k], 1 / _max)
	}
}

func pivot(A [][]float64, i int) {
	_max := A[i][i]
	maxIndex := i

	for k := i + 1; k < len(A); k++ {
		a := A[k][i]

		if a > _max {
			_max = a
			maxIndex = k
		}
	}

	switchLines(A, i, maxIndex)
}

func subtruct(A [][]float64, i int) {
	for j := i + 1; j < len(A); j++ {
		m_ij := A[j][i] / A[i][i]

		addLines(A[j], A[i], -m_ij)
	}
}

func solve(A [][]float64) []float64{
	n := len(A) - 1
	m := len(A[0]) - 1

	X := make([]float64, n + 1)
	
	X[n] = A[n][m] / A[n][n]

	for i := n - 1; i >= 0; i-- {
		sigma := calculateSum_solve(A[i], X, i, n)
		X[i] = (A[i][m] - sigma) / A[i][i]
	}

	return X
}





func getMax(A [][]float64, k int) float64 {
	_max := math.Abs(A[k][0])

	for i, a := range A[k] {
		if i < len(A[k]) - 1 && math.Abs(a) > _max {
			_max = math.Abs(a)
		}
	}

	return _max
}

func multiplieLine(line []float64, coefficient float64) {
	for i, l := range line {
		line[i] = coefficient * l
	}
}

func switchLines(A [][]float64, line1 int, line2 int){
	if line1 == line2 {
		return
	}

	r1 := A[line1]
	r2 := A[line2]

	A[line1] = r2
	A[line2] = r1
}

func addLines(destination []float64, toAdd []float64, coefficient float64){
	if coefficient == 0 {
		return
	}

	toAddMultiplied := []float64{}
	
	// multiplie toAdd by coefficient
	for _, l := range toAdd {
		toAddMultiplied = append(toAddMultiplied, coefficient * l)
	}

	// add toAddMultiplied to destination
	for i, l := range destination {
		destination[i] = l + toAddMultiplied[i]
	}
}

func calculateSum_solve(a_i []float64, X []float64, i int, n int) float64 {
	result := 0.0

	for j := i+1; j <= n; j++ {
		result += a_i[j] * X[j]
	}

	return result
}