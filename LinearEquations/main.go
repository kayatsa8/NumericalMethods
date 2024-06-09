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

	router.Mount("/api/solve_equations", Routes())

	fmt.Println("Starting...")

    http.ListenAndServe(":3004", router)
}

func Routes() chi.Router {
	router := chi.NewRouter()

	controller := NewEquationController()

	router.Post("/gaussian", controller.GaussianEliminationAPI)

	return router
}





