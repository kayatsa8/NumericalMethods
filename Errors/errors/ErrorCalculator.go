package errors

import "math"

func AbsoluteError(original float64, estimated float64) float64 {
	return math.Abs(original-estimated)
}

func RelativeError(original float64, estimated float64) float64 {
	return AbsoluteError(original, estimated) / math.Abs(original)
}

/*
	result < 1 ---> f reduce the error
	result = 1 ---> f preserve the error
	result > 1 ---> f increase the error
*/
func FunctionEffectOnResult(original []float64, estimated []float64, f func([]float64)(float64)) float64{
	var inputRelativeError float64 = sumRelativeError(original, estimated)

	var outputRelativeError float64 = RelativeError(f(original), f(estimated))

	return outputRelativeError / inputRelativeError
}

func sumRelativeError(original []float64, estimated []float64) float64 {
	var sum float64 = 0

	for index := range original {
		sum += RelativeError(original[index], estimated[index])
	}

	return sum
}