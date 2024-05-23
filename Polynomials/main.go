package main

import (
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

    http.ListenAndServe(":3000", router)
}

func Routes() chi.Router {
	router := chi.NewRouter()

	controller := NewPolynomialController()

	router.Get("/add", controller.AddPolynomials)
	router.Get("/addList", controller.AddPolynomialList)
	router.Get("/multiplie", controller.MultipliePolynomials)
	router.Get("/multiplieList", controller.MultipliePolynomialList)
	router.Get("/calculate", controller.Calculate)
	router.Get("/to_string", controller.ToString)

	return router
}