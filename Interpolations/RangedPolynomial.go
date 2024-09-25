package main




type RangedPolynomial struct {
	Pol Polynomial `json:"polynomial"`
	Low float64 `json:"low"`
	Up float64 `json:"up"`
}

func NewRangedPolynomial(p Polynomial, low float64, up float64) RangedPolynomial {
	polynomial := RangedPolynomial{
		Pol: p,
		Low: low,
		Up: up,
	}

	return polynomial
}












