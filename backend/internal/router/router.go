package router

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/handler"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/go-chi/chi/v5"
)

var userRepo = repository.NewUserRepository()
var slamRepo = repository.NewSlamRepository()
var slamPartRepo = repository.NewSlamParticipationRepository()

func SetupV1Router() *chi.Mux {

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

	return r

}

func SetupTestRouter() *chi.Mux {
	r := SetupV1Router()
	authService := service.NewAuthService(userRepo)

	u1, err := authService.Register("Alice", "alice@example.com", "P@ssw0rd")
	if err != nil {
		panic("Error when creating test user 1")
	}
	u2, err := authService.Register("Bob", "bob@example.com", "P@ssw0rd")
	if err != nil {
		panic("Error when creating test user 2")
	}

	slamId := 1
	slamTitle := "Poetry Night"
	slamDescription := "Evening of poems"
	isPublic := true

	slamRepo.Create(api.Slam{
		Id:          &slamId,
		Title:       &slamTitle,
		Description: &slamDescription,
		Public:      &isPublic,
	})
	slamRepo.Create(api.Slam{
		Id:          &slamId,
		Title:       &slamTitle,
		Description: &slamDescription,
		Public:      &isPublic,
	})

	slamPartRepo.Add(*u1.Id, 1)
	slamPartRepo.Add(*u2.Id, 2)

	return r
}
