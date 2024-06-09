package main

import "fmt"

func main() {
	matrix := [][]float64{
		{2, 4, 1, 5},
		{10, 4000, 1, 500},
		{0, -4, 0.01, 1},
	}

	X := GaussianElimination(matrix)

	fmt.Println(X)
}