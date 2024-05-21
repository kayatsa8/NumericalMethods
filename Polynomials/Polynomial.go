package main

import (
	"fmt"
	"math"
	"sort"
)

type Polynomial struct {
	coefficients map[int]float64 // <degree, coefficient>
}


// ----------------- constructors -----------------
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


// ----------------- add polynomials -----------------

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


// ----------------- multiplie polynomials -----------------

func multipliePolynomials(polynomial1 *Polynomial, polynomial2 *Polynomial) (Polynomial, error) {
	if polynomial1 == nil || polynomial2 == nil {
		return emptyPolynomial(), fmt.Errorf("one of the polynomials is nil")
	}

	result := emptyPolynomial()

	for degree1, coefficient1 := range polynomial1.coefficients {
		for degree2, coefficient2 := range polynomial2.coefficients {
			resDegree := degree1 + degree2
			resCoefficient := coefficient1 * coefficient2

			if _, ok := result.coefficients[resDegree]; ok {
				result.coefficients[resDegree] += resCoefficient
			} else{
				result.coefficients[resDegree] = resCoefficient
			}
		}
	}

	return result, nil
}

func multipliePolynomialList(polynomials []Polynomial, coefficient float64) (Polynomial, error) {
	if polynomials == nil {
		return emptyPolynomial(), fmt.Errorf("the polynomial list is nil")
	}

	result := newPolynomial(map[int]float64 {0: coefficient})
	var err error = nil 

	for _, polynomial := range polynomials {
		result, err = multipliePolynomials(&result, &polynomial)

		if err != nil{
			return emptyPolynomial(), err
		}
	}

	return result, nil
}




// ----------------- to string -----------------

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





