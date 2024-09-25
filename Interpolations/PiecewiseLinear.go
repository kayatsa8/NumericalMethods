package main

import (
	"fmt"
	"sort"
)

// gets a map of <x, y> points
func PiecewiseLinear(points map[float64]float64) ([]RangedPolynomial, error){
	Ls := []RangedPolynomial{}

	if len(points) == 0 {
		return Ls, fmt.Errorf("the points collection is empty")
	}

	X := getX(points)

	for i := 0; i<len(X)-1; i++{
		a := (points[X[i + 1]] - points[X[i]]) / (X[i + 1] - X[i])
		b := points[X[i]] - a * X[i]

		p := NewRangedPolynomial(NewPolynomial(map[int]float64{
			0: b,
			1: a,
		}), X[i], X[i + 1])

		Ls = append(Ls, p)
	}
	
	return Ls, nil
}

func getX(points map[float64]float64) []float64 {
	X := make([]float64, len(points))

	i := 0

	for x := range points {
		X[i] = x
	}

	sort.Slice(X, func(i, j int) bool {
		return X[i] < X[j]
	})

	return X
}
















