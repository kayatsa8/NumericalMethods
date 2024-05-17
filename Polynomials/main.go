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
		0: 1,
		6: 91,
	})

	fmt.Println(pol1.toString())
	fmt.Println(pol2.toString())


	newPol, err := pol1.addPolynomial(&pol2)

	fmt.Println(newPol.toString())
	fmt.Println(err)
}