package main

import "fmt"

func main() {
	points := map[float64]float64{
		1.57: 2.63,
		5.98: 4.52,
		11.87: 0.23,
		17.25: 11.2,
		20.21: 0.0,
	}

	pol, err := newtonInterpolation(points)

	fmt.Println(fetch[string]("http://localhost:3000/api/polynomials/to_string", pol).Result)

	fmt.Println(err)
}