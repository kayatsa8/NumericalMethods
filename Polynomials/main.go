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

	router.Mount("/api/polynomials", Routes())

	fmt.Println("Starting...")

    http.ListenAndServe(":3000", router)
}

func Routes() chi.Router {
	router := chi.NewRouter()

	controller := NewPolynomialController()

	router.Post("/add", controller.AddPolynomials)
	router.Post("/multiplie", controller.MultipliePolynomials)
	router.Post("/calculate", controller.Calculate)
	router.Post("/to_string", controller.ToString)

	return router
}