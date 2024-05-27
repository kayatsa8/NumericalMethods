package main

import (
	"encoding/json"
	"net/http"

	"ErrorCalculator/errors"
)

type ErrorController struct {
}

func NewErrorController() ErrorController {
	return ErrorController{}
}

func (controller ErrorController) AbsoluteError(w http.ResponseWriter, r *http.Request) {
	var input ErrorInput

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	result := errors.AbsoluteError(input.Original, input.Estimated)
	response := NewResponse(result, nil)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller ErrorController) RelativeError(w http.ResponseWriter, r *http.Request) {
	var input ErrorInput

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	result := errors.RelativeError(input.Original, input.Estimated)
	response := NewResponse(result, nil)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller ErrorController) FunctionEffectOnResult(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("The method is not implemented yet!")
	checkInternalError(err, w)
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
