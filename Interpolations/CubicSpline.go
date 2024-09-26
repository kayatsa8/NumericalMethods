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

	a := getA(points, X)
	h := getH(X)
	c := getC(h, a)
	b := getB(h, a, c)
	d := getD(c, h)

	for i:=0; i<len(X)-1; i++ {
		s1 := NewPolynomial(map[int]float64{
			0: a[i],
		})
		
		sTemp := NewPolynomial(map[int]float64{
			0: -X[i],
			1: 1,
		})

		s2 := multipliePolynomials([]Polynomial{sTemp}, b[i])
		s3 := multipliePolynomials([]Polynomial{sTemp, sTemp}, c[i])
		s4 := multipliePolynomials([]Polynomial{sTemp, sTemp, sTemp}, d[i])

		s := addMultiplePolynomials([]Polynomial{s1, s2, s3, s4})

		S = append(S, NewRangedPolynomial(s, X[i], X[i + 1]))
	}
	
	return S, nil
}

func getA(points map[float64]float64, X []float64) []float64 {
	a := make([]float64, len(X) - 1)

	for i := 0; i<len(X); i++ {
		a[i] = points[X[i]]
	}

	return a
}

func getH(X []float64) []float64 {
	h := make([]float64, len(X) - 1)

	for i:=0; i<len(X)-1; i++{
		h[i] = X[i + 1] - X[i]
	}

	return h
}

func getC(h []float64, a []float64) []float64 {
	n := len(h)
	equations := make([][]float64, n - 1, n)

	equations[0][0] = 2 * (h[0] + h[1])
	equations[0][1] = h[1]
	equations[0][n - 1] = (3 / h[1]) * (a[2] - a[1]) - (3 / a[0]) * (a[1] - a[0])

	for i:=1; i<n-1; i++ {
		equations[i][i - 1] = h[i]
		equations[i][i] = 2 * (h[i] + h[i + 1])
		equations[i][i + 1] = h[i + 1]

		equations[i][n - 1] = (3 / h[i + 1]) * (a[i + 2] - a[i + 1]) - (3 / h[i]) * (a[i + 1] - a[i])
	}

	equations[n - 2][n - 3] = h[n - 2]
	equations[n - 2][n - 2] = 2 * (h[n - 2] + h[n - 1])
	equations[n - 2][n - 1] = (3 / h[n]) * (a[n + 1] - a[n]) - (3 / h[n - 1]) * (a[n] - a[n - 1])
	
	c := solveEquations(equations)
	c = append([]float64{0}, c...)
	c = append(c, 0.0)
	
	return c
}

func getB(h []float64, a []float64, c []float64) []float64 {
	n := len(a) - 1
	b := make([]float64, n)

	for i:=0; i<n; i++ {
		b[i] = (1 / h[i]) * (a[i + 1] - a[i]) - (h[i] / 3) * (2 * c[i] + c[i + 1])
	}

	return b
}

func getD(c []float64, h []float64) []float64 {
	n := len(h)
	d := make([]float64, n)

	for i:=0; i<n; i++{
		d[i] = (c[i + 1] - c[i]) / (3 * h[i])
	}

	return d
}




func solveEquations(equations [][]float64) []float64{
	response := fetch[[]float64]("http://localhost:3004/api/solve_equations", equations)

	if response.Err != nil {
		log.Fatal(response.Err)
		return []float64{}
	}

	return response.Result
}












