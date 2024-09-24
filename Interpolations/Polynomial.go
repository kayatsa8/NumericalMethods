package main

import (
	"log"
)


type Polynomial struct {
	Coefficients map[int]float64 `json:"coefficients"` // <degree, coefficient>
}


// ----------------- constructors -----------------
func NewPolynomial(coefficients map[int]float64) Polynomial {
	polynomial := Polynomial{
		Coefficients: coefficients,
	}

	polynomial.removeZeroCoefficients()

	return polynomial
}

func EmptyPolynomial() Polynomial{
	return Polynomial{map[int]float64{}}
}

func (polynomial *Polynomial) removeZeroCoefficients() {
	for degree, coefficient := range polynomial.Coefficients {
		if coefficient == 0 {
			delete(polynomial.Coefficients, degree)
		}
	}

	if len(polynomial.Coefficients) == 0 {
		polynomial.Coefficients[0] = 0
	}
}


func multipliePolynomials(polynomials []Polynomial, coefficients float64) Polynomial {
	response := fetch[Polynomial]("http://localhost:3000/api/polynomials/multiplie",
						MultiplieInput{Polynomianls: polynomials, Coefficient: coefficients})

	if response.Err != nil {
		log.Fatal(response.Err)
		return EmptyPolynomial()
	}

	return response.Result
}

func addPolynomials(pol1 *Polynomial, pol2 *Polynomial) Polynomial {
	response := fetch[Polynomial]("http://localhost:3000/api/polynomials/add", []Polynomial{*pol1, *pol2})

	if response.Err != nil {
		log.Fatal(response.Err)
		return EmptyPolynomial()
	}

	return response.Result
}





