package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/handler"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// testData := flag.Bool("test-data", false, "Start the server instance with some test data")
	flag.Parse()

	userService := service.NewUserService(repository.NewUserRepository())
	slamService := service.NewSlamService(repository.NewSlamRepository())
	authService := service.NewAuthService(repository.NewUserRepository())
	partService := service.NewSlamParticipationService(repository.NewUserRepository(), repository.NewSlamRepository(), repository.NewSlamParticipationRepository())

	r := chi.NewRouter()
	server := handler.NewServer(userService, slamService, authService, partService)
	// server := api.Unimplemented{}

	spec, err := api.LoadSpec()
	if err != nil {
		panic(err)
	}

	r.Route("/api/v1", func(apiV1 chi.Router) {
		apiV1.Use(utils.AuthMiddleware(spec))
		api.HandlerFromMux(server, apiV1)
	})

	r.Get(api.SpecUrl, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, api.SpecPath)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(api.SpecUrl),
	))

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Println("Server starting on :8080")
	log.Fatal(s.ListenAndServe())
}
