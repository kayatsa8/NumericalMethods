package main


type BisectionRegulaInput struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	Delta float64 `json:"delta"`
	F string `json:"f"`
}