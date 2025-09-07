package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Mroxny/slamIt/internal/handler"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/service"
)

func main() {
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)
	slamRepo := repository.NewSlamRepository()
	slamService := service.NewSlamService(slamRepo)
	slamHandler := handler.NewSlamHandler(slamService)
	partRepo := repository.NewSlamParticipationRepository()
	partService := service.NewSlamParticipationService(repo, slamRepo, partRepo)
	partHandler := handler.NewSlamParticipationHandler(partService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)
		r.Post("/", userHandler.Create)
		r.Get("/{id}", userHandler.GetByID)
		r.Put("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)
	})

	r.Route("/slams", func(r chi.Router) {
		r.Get("/", slamHandler.GetAll)
		r.Post("/", slamHandler.Create)
		r.Get("/{id}", slamHandler.GetByID)
		r.Put("/{id}", slamHandler.Update)
		r.Delete("/{id}", slamHandler.Delete)
	})

	r.Route("/participation", func(r chi.Router) {
		r.Post("/users/{userID}/slams/{slamID}", partHandler.Join)
		r.Delete("/users/{userID}/slams/{slamID}", partHandler.Leave)
		r.Get("/users/{userID}/slams", partHandler.GetSlamsForUser)
		r.Get("/slams/{slamID}/users", partHandler.GetUsersForSlam)
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
