package main

import (
	"fmt"
	"math"
)

func main() {

	root := fixedPoint(1, func(x float64) float64{return math.Pow(3 + x - 2 * math.Pow(x, 2), 0.25)}, 0.0001)

	fmt.Printf("The root is: %v", root)
}