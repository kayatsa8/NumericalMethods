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

	router.Mount("/api/roots", Routes())

	fmt.Println("Starting...")

    http.ListenAndServe(":3002", router)
}

func Routes() chi.Router {
	router := chi.NewRouter()

	controller := NewRootController()

	router.Get("/bisection", controller.Bisection)

	return router
}