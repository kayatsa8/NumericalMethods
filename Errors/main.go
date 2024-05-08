package main

import (
	"ErrorCalculator/errors"
	"fmt"
)

func main() {
	x := 111116.5891
	x_tilda := 111116.59

	y := 35613.11235
	y_tilda := 35613.11


	fmt.Printf("Absolute error x: %v\n", errors.AbsoluteError(x, x_tilda))
	fmt.Printf("Relative error x: %v\n", errors.RelativeError(x, x_tilda))

	fmt.Printf("Absolute error y: %v\n", errors.AbsoluteError(y, y_tilda))
	fmt.Printf("Relative error y: %v\n", errors.RelativeError(y, y_tilda))

	fmt.Printf("Function effect: %v\n", errors.FunctionEffectOnResult([]float64{x, y}, []float64{x_tilda, y_tilda}, mult))
}

func sum(x []float64) float64 {
	var sum float64 = 0

	for _, value := range x {
		sum += value
	}

	return sum
}

func mult(x []float64) float64 {
	var res float64 = 1

	for _, value := range x {
		res *= value
	}

	return res
}