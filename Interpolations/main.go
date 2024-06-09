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

	router.Mount("/api/interpolation", Routes())

	fmt.Println("Starting...")

    http.ListenAndServe(":3003", router)
}

func Routes() chi.Router {
	router := chi.NewRouter()

	controller := NewInterpolationController()

	router.Post("/newton", controller.NewtonInterpolation)

	return router
}