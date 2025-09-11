package main

import (
	"log"
	"net/http"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/api/v1", router.SetupV1Router())
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
