package main

import (
	"encoding/json"
	"net/http"
)

type EquationController struct {
}

func NewEquationController() EquationController {
	return EquationController{}
}


func (controller EquationController) GaussianEliminationAPI(w http.ResponseWriter, r *http.Request) {
	var Equations [][]float64

	err0 := json.NewDecoder(r.Body).Decode(&Equations)
	checkBadRequest(err0, w)

	result, err := GaussianElimination(Equations)
	response := NewResponse(result, err)

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