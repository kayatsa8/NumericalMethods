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

	m1, _ := multipliePolynomials(&pol1, &pol2)
	m2, _ := multipliePolynomials(&m1, &pol3)
	lm, _ := multipliePolynomialList([]Polynomial{pol1, pol2, pol3}, 1)

	fmt.Println(m2.toString())
	fmt.Println(lm.toString())
}