package main

import "net/http"

type PolynomialController struct {
}


func NewPolynomialController() PolynomialController {
	return PolynomialController{}
}


func (controller PolynomialController) AddPolynomials(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("AddPolynomails"))
}

func (controller PolynomialController) AddPolynomialList(w http.ResponseWriter, r *http.Request){}

func (controller PolynomialController) MultipliePolynomials(w http.ResponseWriter, r *http.Request){}

func (controller PolynomialController) MultipliePolynomialList(w http.ResponseWriter, r *http.Request){}

func (controller PolynomialController) Calculate(w http.ResponseWriter, r *http.Request){}

func (controller PolynomialController) ToString(w http.ResponseWriter, r *http.Request){}