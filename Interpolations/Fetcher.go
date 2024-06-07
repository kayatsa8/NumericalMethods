package main

import(
	"encoding/json"
    "net/http"
	"bytes"
)

func fetch[T any](url string, data any) Response[T] {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return Response[T]{Err: err,}
	}

	apiResponse, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return Response[T]{Err: err,}
    }

	var response Response[T]
	err = json.NewDecoder(apiResponse.Body).Decode(&response)

	if err != nil {
		return Response[T]{Err: err,}
	}

	return response
}