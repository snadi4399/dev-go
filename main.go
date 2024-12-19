package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/cities", listCities)
	r.Get("/cities/{city_id}", viewCity)
	r.Get("/cities/{city_id}/edit", editCity)
	r.Post("/cities/{city_id}/edit", saveCity)

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", r)
}
