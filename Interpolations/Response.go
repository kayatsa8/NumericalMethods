package main

type Response[T any] struct {
	Result T `json:"result"`
	Err error `json:"error"`
}

func NewResponse[T any](result T, err error) Response[T] {
	return Response[T]{
		Result: result,
		Err: err,
	}
}