package main

import (
	"encoding/json"
	"net/http"

    "github.com/Knetic/govaluate"
)

type RootController struct{}

func NewRootController() RootController{
	return RootController{}
}

func (controller RootController) Bisection(w http.ResponseWriter, r *http.Request) {
	var input BisectionRegulaInput

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	f := makeFunction(input.F, w)

	result, err := bisection(input.A, input.B, f, input.Delta)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}









func makeFunction(funcString string, w http.ResponseWriter) func(float64)float64 {
	expression, err := govaluate.NewEvaluableExpression(funcString)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
    }

	return func(x float64)float64{
		parameters := map[string]interface{}{
			"x": 2,
		}

		result, _ := expression.Evaluate(parameters)
		
		return result.(float64)
	}
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