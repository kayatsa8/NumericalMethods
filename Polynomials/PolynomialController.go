package main

import (
	"encoding/json"
	"net/http"
)

type PolynomialController struct {
}

func NewPolynomialController() PolynomialController {
	return PolynomialController{}
}

func (controller PolynomialController) AddPolynomials(w http.ResponseWriter, r *http.Request) {
	var polynomials []Polynomial

	err0 := json.NewDecoder(r.Body).Decode(&polynomials)
	checkBadRequest(err0, w)

	result, err := addAllPolynomials(polynomials)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller PolynomialController) MultipliePolynomials(w http.ResponseWriter, r *http.Request) {
	var input multiplieInput

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	result, err := multipliePolynomialList(input.Polynomianls, input.Coefficient)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller PolynomialController) Calculate(w http.ResponseWriter, r *http.Request) {
	var input calculateInput

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	result, err := input.Pol.calculate(input.Value)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller PolynomialController) Derivative(w http.ResponseWriter, r *http.Request){
	var input Polynomial

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	result := input.derivative()
	response := NewResponse(result, nil)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller PolynomialController) ToString(w http.ResponseWriter, r *http.Request) {
	var polynomial Polynomial

	err0 := json.NewDecoder(r.Body).Decode(&polynomial)
	checkBadRequest(err0, w)

	result := polynomial.toString()
	response := NewResponse(result, nil)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func checkInternalError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func checkBadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
