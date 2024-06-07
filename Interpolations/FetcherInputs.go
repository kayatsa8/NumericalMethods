package main

type calculateInput struct {
	Pol Polynomial `json:"polynomial"`
	Value float64 `json:"value"`
}

type multiplieInput struct {
	Polynomianls []Polynomial `json:"polynomials"`
	Coefficient float64 `json:"coefficient"`
}
