package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type InterpolationController struct {
}

func NewInterpolationController() InterpolationController {
	return InterpolationController{}
}

func (controller InterpolationController) NewtonInterpolation(w http.ResponseWriter, r *http.Request) {
	var pointsString map[string]float64

	err0 := json.NewDecoder(r.Body).Decode(&pointsString)
	checkBadRequest(err0, w)

	points := makePointsMap(pointsString, w)

	result, err := NewtonInterpolation(points)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller InterpolationController) VandermondeInterpolation(w http.ResponseWriter, r *http.Request) {
	var pointsString map[string]float64

	err0 := json.NewDecoder(r.Body).Decode(&pointsString)
	checkBadRequest(err0, w)

	points := makePointsMap(pointsString, w)

	result, err := VandermondeInterpolation(points)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller InterpolationController) LagrangeInterpolation(w http.ResponseWriter, r *http.Request) {
	var pointsString map[string]float64

	err0 := json.NewDecoder(r.Body).Decode(&pointsString)
	checkBadRequest(err0, w)

	points := makePointsMap(pointsString, w)

	result, err := LagrangeInterpolation(points)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller InterpolationController) HermiteInterpolation(w http.ResponseWriter, r *http.Request) {
	var input map[string]map[string]float64

	err0 := json.NewDecoder(r.Body).Decode(&input)
	checkBadRequest(err0, w)

	points := makePointsMap(input["points"], w)
	diffPoints := makePointsMap(input["diffPoints"], w)

	result, err := HermiteInterpolation(points, diffPoints)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller InterpolationController) PiecewiseLinear(w http.ResponseWriter, r *http.Request) {
	var pointsString map[string]float64

	err0 := json.NewDecoder(r.Body).Decode(&pointsString)
	checkBadRequest(err0, w)

	points := makePointsMap(pointsString, w)

	result, err := PiecewiseLinear(points)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}

func (controller InterpolationController) CubicSpline(w http.ResponseWriter, r *http.Request) {
	var pointsString map[string]float64

	err0 := json.NewDecoder(r.Body).Decode(&pointsString)
	checkBadRequest(err0, w)

	points := makePointsMap(pointsString, w)

	result, err := CubicSpline(points)
	response := NewResponse(result, err)

	err1 := json.NewEncoder(w).Encode(response)
	checkInternalError(err1, w)
}


func makePointsMap(pointsString map[string]float64, w http.ResponseWriter) map[float64]float64{
	points := map[float64]float64{}

	for key, value := range pointsString {
		floatKey, err := strconv.ParseFloat(key, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return map[float64]float64{}
		}
		points[floatKey] = value
	}

	return points;
}













func checkInternalError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

func checkBadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}