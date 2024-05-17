package main

import "fmt"

func main() {
	pol1 := newPolynomial(map[int]float64{
		2: 5,
		9: 22,
		0: 58,
		1: -8,
		7: 0,
	})

	fmt.Println(pol1.toString())
}