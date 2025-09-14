package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Mroxny/slamIt/internal/handler"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/service"
)

var userRepo = repository.NewUserRepository()
var slamRepo = repository.NewSlamRepository()
var partRepo = repository.NewSlamParticipationRepository()

func SetupV1Router() http.Handler {

	// --- services ---
	userService := service.NewUserService(userRepo)
	slamService := service.NewSlamService(slamRepo)
	partService := service.NewSlamParticipationService(userRepo, slamRepo, partRepo)
	authService := service.NewAuthService(userRepo)

	// --- handlers ---
	userHandler := handler.NewUserHandler(userService)
	slamHandler := handler.NewSlamHandler(slamService)
	partHandler := handler.NewSlamParticipationHandler(partService)
	authHandler := handler.NewAuthHandler(authService)

	// --- router ---
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// --- auth routes ---
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	// --- public slam routes ---
	r.Route("/slams", func(r chi.Router) {
		r.Get("/", slamHandler.GetAll)

		r.Group(func(r chi.Router) {
			r.Use(handler.AuthMiddleware())

			r.Post("/", slamHandler.Create)
			r.Get("/{id}", slamHandler.GetByID)
			r.Put("/{id}", slamHandler.Update)
			r.Delete("/{id}", slamHandler.Delete)
		})
	})

	// --- users (protected) ---
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)

		r.Group(func(r chi.Router) {
			r.Use(handler.AuthMiddleware())
			r.Get("/{id}", userHandler.GetByID)
			r.Put("/{id}", userHandler.Update)
			r.Delete("/{id}", userHandler.Delete)
		})
	})

	// --- participation (protected) ---
	r.Route("/participation", func(r chi.Router) {
		r.Use(handler.AuthMiddleware())
		r.Post("/users/{userID}/slams/{slamID}", partHandler.Join)
		r.Delete("/users/{userID}/slams/{slamID}", partHandler.Leave)
		r.Get("/users/{userID}/slams", partHandler.GetSlamsForUser)
		r.Get("/slams/{slamID}/users", partHandler.GetUsersForSlam)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})

	return r
}

func SetupTestRouter() http.Handler {
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

	slamRepo.Create(model.Slam{
		ID:          1,
		Title:       "Poetry Night",
		Description: "Evening of poems",
		Public:      true,
	})
	slamRepo.Create(model.Slam{
		ID:          2,
		Title:       "Secret Slam",
		Description: "Invite only",
		Public:      false,
	})

	partRepo.Add(u1.ID, 1)
	partRepo.Add(u2.ID, 2)

	return r
}
