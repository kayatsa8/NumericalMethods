package main

type multiplieInput struct {
	Polynomianls []Polynomial `json:"polynomials"`
	Coefficient float64 `json:"coefficient"`
}

type calculateInput struct {
	Pol Polynomial `json:"polynomial"`
	Value float64 `json:"value"`
}