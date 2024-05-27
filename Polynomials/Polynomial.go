package main

import (
	"fmt"
	"math"
	"sort"
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


// ----------------- add polynomials -----------------

func (polynomial *Polynomial) addPolynomial(toAdd *Polynomial) (Polynomial, error) {
	if toAdd == nil {
		return Polynomial{}, fmt.Errorf("given Polynomial is nil")
	}

	newPol := Polynomial{map[int]float64{}}

	// add same-degree coefficients and unique-degree coefficients of polynomial
	for degree, coefficient := range polynomial.Coefficients {
		if coefficient2, ok := toAdd.Coefficients[degree]; ok {
			newPol.Coefficients[degree] = coefficient + coefficient2
		} else {
			newPol.Coefficients[degree] = coefficient
		}
	}

	// add unique-degree coefficients of toAdd
	for degree, coefficient := range toAdd.Coefficients {
		if _, ok := polynomial.Coefficients[degree]; !ok {
			newPol.Coefficients[degree] = coefficient
		}
	}

	newPol.removeZeroCoefficients()

	return newPol, nil
}

func addAllPolynomials(polynomials []Polynomial) (Polynomial, error) {
	result := EmptyPolynomial()
	err := fmt.Errorf("")

	for _, polynomial := range polynomials {
		result, err = result.addPolynomial(&polynomial)

		if err != nil {
			return EmptyPolynomial(), err
		}
	}

	return result, nil
}


// ----------------- multiplie polynomials -----------------

func multipliePolynomials(polynomial1 *Polynomial, polynomial2 *Polynomial) (Polynomial, error) {
	if polynomial1 == nil || polynomial2 == nil {
		return EmptyPolynomial(), fmt.Errorf("one of the polynomials is nil")
	}

	result := EmptyPolynomial()

	for degree1, coefficient1 := range polynomial1.Coefficients {
		for degree2, coefficient2 := range polynomial2.Coefficients {
			resDegree := degree1 + degree2
			resCoefficient := coefficient1 * coefficient2

			if _, ok := result.Coefficients[resDegree]; ok {
				result.Coefficients[resDegree] += resCoefficient
			} else{
				result.Coefficients[resDegree] = resCoefficient
			}
		}
	}

	return result, nil
}

func multipliePolynomialList(polynomials []Polynomial, coefficient float64) (Polynomial, error) {
	if polynomials == nil {
		return EmptyPolynomial(), fmt.Errorf("the polynomial list is nil")
	}

	result := NewPolynomial(map[int]float64 {0: coefficient})
	var err error = nil 

	for _, polynomial := range polynomials {
		result, err = multipliePolynomials(&result, &polynomial)

		if err != nil{
			return EmptyPolynomial(), err
		}
	}

	return result, nil
}


// ----------------- calculate -----------------

func (polynomial *Polynomial) calculate(value float64) (float64, error) {
	result := 0.0

	for degree, coefficient := range polynomial.Coefficients {
		result += coefficient * math.Pow(value, float64(degree))
	}

	return result, nil
}



// ----------------- to string -----------------

func (polynomial *Polynomial) toString() string {
	pol := "P(x) ="
	first := true

	degrees := sortDegreesToString(polynomial.Coefficients)

	for _, degree := range degrees {
		pol = addCoefficientToString(pol, polynomial.Coefficients[degree], degree, first)
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





