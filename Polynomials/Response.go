package main

type Response struct {
	Result any `json:"result"`
	Err error `json:"error"`
}

func NewResponse(result any, err error) Response {
	return Response{
		Result: result,
		Err: err,
	}
}