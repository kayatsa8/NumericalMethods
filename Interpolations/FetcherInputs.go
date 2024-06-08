package main


type CalculateInput struct {
	Pol Polynomial `json:"polynomial"`
	Value float64 `json:"value"`
}

type MultiplieInput struct {
	Polynomianls []Polynomial `json:"polynomials"`
	Coefficient float64 `json:"coefficient"`
}
