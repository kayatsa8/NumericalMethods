package main

import (
	"fmt"
	"log"
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
	C := getC(H, A)
	
	return S, nil
}

func getA(points map[float64]float64, X []float64) []float64 {
	A := make([]float64, len(X) - 1)

	for i := 0; i<len(X); i++ {
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
	n := len(H)
	equations := make([][]float64, n - 1, n)

	equations[0][0] = 2 * (H[0] + H[1])
	equations[0][1] = H[1]
	equations[0][n - 1] = (3 / H[1]) * (A[2] - A[1]) - (3 / H[0]) * (A[1] - A[0])

	for i:=1; i<n-1; i++ {
		equations[i][i - 1] = H[i]
		equations[i][i] = 2 * (H[i] + H[i + 1])
		equations[i][i + 1] = H[i + 1]

		equations[i][n - 1] = (3 / H[i + 1]) * (A[i + 2] - A[i + 1]) - (3 / H[i]) * (A[i + 1] - A[i])
	}

	equations[n - 2][n - 3] = H[n - 2]
	equations[n - 2][n - 2] = 2 * (H[n - 2] + H[n - 1])
	equations[n - 2][n - 1] = (3 / H[n]) * (A[n + 1] - A[n]) - (3 / H[n - 1]) * (A[n] - A[n - 1])
	
	C := solveEquations(equations)
	C = append([]float64{0}, C...)
	C = append(C, 0.0)
	
	return C
}

func solveEquations(equations [][]float64) []float64{
	response := fetch[[]float64]("http://localhost:3004/api/solve_equations", equations)

	if response.Err != nil {
		log.Fatal(response.Err)
		return []float64{}
	}

	return response.Result
}










