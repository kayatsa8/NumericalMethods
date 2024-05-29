package main


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

