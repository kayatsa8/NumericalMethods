package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

	router.Mount("/api/errors", Routes())

	fmt.Println("Starting...")

    http.ListenAndServe(":3001", router)
}

func Routes() chi.Router {
	router := chi.NewRouter()

	controller := NewErrorController()

	router.Get("/absolute_error", controller.AbsoluteError)
	router.Get("/relative_error", controller.RelativeError)
	router.Get("/effect_on_result", controller.FunctionEffectOnResult)

	return router
}