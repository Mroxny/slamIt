package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/config"
	"github.com/Mroxny/slamIt/internal/router"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg := config.GetConfig()
	testData := flag.Bool("test-data", false, "Start the server instance with some test data")
	flag.Parse()

	var r *chi.Mux
	if *testData {
		r = router.SetupTestRouter()
	} else {
		r = router.SetupV1Router()
	}

	r.Get(api.SpecUrl, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, api.SpecPath)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(api.SpecUrl),
	))

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + cfg.Port,
	}

	log.Println("Server starting on :" + cfg.Port)
	log.Fatal(s.ListenAndServe())
}
