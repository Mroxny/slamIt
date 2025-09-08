package main

import (
	"log"
	"net/http"

	"github.com/Mroxny/slamIt/internal/router"
)

func main() {
	r := router.SetupRouter()
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
