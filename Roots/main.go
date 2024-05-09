package main

import (
	"fmt"
	"math"
)

func main() {

	root, err := regulaFalsi(5, 6, func(x float64) float64{return (math.Exp(x) / math.Pow(x, 2)) - 10}, 0.001)

	if err == nil {
		fmt.Printf("The root is: %v", root)
	} else{
		fmt.Println(err)
	}

	
}