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

	pol2 := newPolynomial(map[int]float64{
		2: -3,
		0: -1,
		6: -1,
	})

	pol3 := newPolynomial(map[int]float64{
		9: -20,
		7: 1,
		6: 2,
	})

	fmt.Println(pol1.toString())
	fmt.Println(pol2.toString())
	fmt.Println(pol3.toString())


	newPol, err := addAllPolynomials([]Polynomial{pol1, pol2, pol3})

	fmt.Println(newPol.toString())
	fmt.Println(err)
}