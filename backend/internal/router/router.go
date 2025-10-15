package router

import (
	"context"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/handler"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/go-chi/chi/v5"
)

type Repositories struct {
	userRepo  *repository.UserRepository
	slamRepo  *repository.SlamRepository
	partRepo  *repository.ParticipationRepository
	stageRepo *repository.StageRepository
	perfRepo  *repository.PerformanceRepository
	voteRepo  *repository.VoteRepository
}

func SetupV1Router(localOnly bool) (*chi.Mux, *Repositories) {
	db := repository.InitDB(localOnly)
	repos := Repositories{
		userRepo:  repository.NewUserRepository(db),
		slamRepo:  repository.NewSlamRepository(db),
		partRepo:  repository.NewParticipationRepository(db),
		stageRepo: repository.NewStageRepository(db),
		perfRepo:  repository.NewPerformanceRepository(db),
		voteRepo:  repository.NewVoteRepository(db),
	}

	userService := service.NewUserService(repos.userRepo)
	slamService := service.NewSlamService(repos.slamRepo)
	authService := service.NewAuthService(repos.userRepo)
	partService := service.NewParticipationService(repos.partRepo)
	stageService := service.NewStageService(repos.stageRepo)
	perfService := service.NewPerformanceService(repos.perfRepo)
	voteService := service.NewVoteService(repos.voteRepo)

	r := chi.NewRouter()
	server := handler.NewServer(
		userService,
		slamService,
		authService,
		partService,
		stageService,
		perfService,
		voteService,
	)
	// server := api.Unimplemented{}

	spec, err := api.LoadSpec()
	if err != nil {
		panic(err)
	}

	r.Route("/api/v1", func(apiV1 chi.Router) {
		apiV1.Use(utils.AuthMiddleware(spec))
		api.HandlerFromMux(server, apiV1)
	})

	return r, &repos

}

func SetupTestRouter() *chi.Mux {
	repository.ClearLocalDB()
	r, repos := SetupV1Router(true)
	ctx := context.Background()
	authService := service.NewAuthService(repos.userRepo)

	u1, err := authService.Register(ctx, "Bob", "bob@example.com", "P@ssw0rd")
	if err != nil {
		panic("Error when creating test user 1")
	}
	u2, err := authService.Register(ctx, "Alice", "alice@example.com", "P@ssw0rd")
	if err != nil {
		panic("Error when creating test user 2")
	}

	slamTitle := "Poetry Night"
	slamDescription := "Evening of poems"

	slam1 := model.Slam{
		Slam: api.Slam{
			Id:          "1b338aa8-74a1-43e9-8034-94f144e77c3a",
			Title:       slamTitle,
			Description: &slamDescription,
			Public:      true,
		},
	}

	slam2 := model.Slam{
		Slam: api.Slam{
			Id:          "85bf4f72-3cd2-46df-8d37-016442f150f7",
			Title:       slamTitle + " 2",
			Description: &slamDescription,
			Public:      false,
		},
	}

	err = repos.slamRepo.CreateWithCreatorTx(ctx, &slam1, u1.Id)
	if err != nil {
		panic("Error when creating test slam 1")
	}
	err = repos.slamRepo.CreateWithCreatorTx(ctx, &slam2, u2.Id)
	if err != nil {
		panic("Error when creating test slam 2")
	}

	return r
}
