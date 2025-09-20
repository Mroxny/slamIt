package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/config"
	"github.com/Mroxny/slamIt/internal/handler"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// testData := flag.Bool("test-data", false, "Start the server instance with some test data")
	flag.Parse()

	userRepo := repository.NewUserRepository()
	slamRepo := repository.NewSlamRepository()
	slamPartRepo := repository.NewSlamParticipationRepository()

	userService := service.NewUserService(userRepo)
	slamService := service.NewSlamService(slamRepo)
	authService := service.NewAuthService(userRepo)
	partService := service.NewSlamParticipationService(userRepo, slamRepo, slamPartRepo)

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
		Addr:    "0.0.0.0:" + cfg.Port,
	}

	log.Println("Server starting on :" + cfg.Port)
	log.Fatal(s.ListenAndServe())
}
