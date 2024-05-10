package main

import (
	"fmt"
	"math"
)

func main() {

	root := newtonMethod(-8, func(x float64) float64{return math.Exp(x) + math.Pow(x, 3) -18*x +7},
						 func(x float64) float64{return math.Exp(x) + 3*math.Pow(x, 2) - 18}, 0.001)

	fmt.Printf("The root is: %v", root)
}