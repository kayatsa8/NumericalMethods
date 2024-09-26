package main

import (
	"fmt"
	"sort"
)


// gets a map of <x, y> points
func CubicSpline(points map[float64]float64) ([]RangedPolynomial, error){
	S := []RangedPolynomial{}

	if len(points) == 0 {
		return S, fmt.Errorf("the points collection is empty")
	}

	X := getX(points)

	A := getA(points, X)
	H := getH(X)
	
	return S, nil
}

func getA(points map[float64]float64, X []float64) []float64 {
	A := make([]float64, len(X) - 1)

	for i := 0; i<len(X)-1; i++ {
		A[i] = points[X[i]]
	}

	return A
}

func getH(X []float64) []float64 {
	H := make([]float64, len(X) - 1)

	for i:=0; i<len(X)-1; i++{
		H[i] = X[i + 1] - X[i]
	}

	return H
}

func getC(H []float64, A []float64) []float64 {
	
}










