package main

import (
	"fmt"
	"math"
	"sort"
)

type Polynomial struct {
	coefficients map[int]float64 // <degree, coefficient>
}

func newPolynomial(coefficients map[int]float64) Polynomial {
	polynomial := Polynomial{
		coefficients: coefficients,
	}

	polynomial.removeZeroCoefficients()

	return polynomial
}

func emptyPolynomial() Polynomial{
	return Polynomial{map[int]float64{}}
}

func (polynomial *Polynomial) removeZeroCoefficients() {
	for degree, coefficient := range polynomial.coefficients {
		if coefficient == 0 {
			delete(polynomial.coefficients, degree)
		}
	}

	if len(polynomial.coefficients) == 0 {
		polynomial.coefficients[0] = 0
	}
}

func (polynomial *Polynomial) addPolynomial(toAdd *Polynomial) (Polynomial, error) {
	if toAdd == nil {
		return Polynomial{}, fmt.Errorf("given Polynomial is nil")
	}

	newPol := Polynomial{map[int]float64{}}

	// add same-degree coefficients and unique-degree coefficients of polynomial
	for degree, coefficient := range polynomial.coefficients {
		if coefficient2, ok := toAdd.coefficients[degree]; ok {
			newPol.coefficients[degree] = coefficient + coefficient2
		} else {
			newPol.coefficients[degree] = coefficient
		}
	}

	// add unique-degree coefficients of toAdd
	for degree, coefficient := range toAdd.coefficients {
		if _, ok := polynomial.coefficients[degree]; !ok {
			newPol.coefficients[degree] = coefficient
		}
	}

	newPol.removeZeroCoefficients()

	return newPol, nil
}

func addAllPolynomials(polynomials []Polynomial) (Polynomial, error) {
	result := emptyPolynomial()
	err := fmt.Errorf("")

	for _, polynomial := range polynomials {
		result, err = result.addPolynomial(&polynomial)

		if err != nil {
			return emptyPolynomial(), err
		}
	}

	return result, nil
}




func (polynomial *Polynomial) toString() string {
	pol := "P(x) ="
	first := true

	degrees := sortDegreesToString(polynomial.coefficients)

	for _, degree := range degrees {
		pol = addCoefficientToString(pol, polynomial.coefficients[degree], degree, first)
		pol = addDegreeToString(pol, degree)
		first = false		
	}

	return pol
}

func sortDegreesToString(coefficients map[int] float64) []int {
	degrees := []int{}

	for degree := range coefficients {
		degrees = append(degrees, degree)
	}

	sort.Ints(degrees)

	return degrees
}

func addCoefficientToString(pol string, coefficient float64, degree int, first bool) string {
	if !first {

		if coefficient > 0 {
			if coefficient == 1 && degree != 0{
				pol = fmt.Sprintf("%v + ", pol)
			} else{
				pol = fmt.Sprintf("%v + %v", pol, coefficient)
			}

			
		} else if coefficient < 0 {
			if coefficient == -1 && degree != 0 {
				pol = fmt.Sprintf("%v - ", pol)
			} else{
				pol = fmt.Sprintf("%v - %v", pol, math.Abs(coefficient))
			}
		}
		
	} else{

		if coefficient >= 0 {
			if coefficient == 1 && degree != 0 {
				pol = fmt.Sprintf("%v ", pol)	
			} else{
				pol = fmt.Sprintf("%v %v", pol, coefficient)
			}
		} else{
			if coefficient == -1 && degree != 0 {
				pol = fmt.Sprintf("%v -", pol)	
			} else{
				pol = fmt.Sprintf("%v -%v", pol, math.Abs(coefficient))
			}
		}

	}

	return pol
}

func addDegreeToString(pol string, degree int) string {
	if degree == 0 {
		return pol
	} else if degree == 1 {
		pol = fmt.Sprintf("%vx", pol)
	} else{
		pol = fmt.Sprintf("%vx^%v", pol, degree)
	}

	return pol;
}





