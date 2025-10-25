package router

import (
	"context"
	"strconv"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/handler"
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

	// USERS
	for i := 0; i < len(utils.TestUsers); i++ {
		if err := repos.userRepo.Create(ctx, &utils.TestUsers[i]); err != nil {
			panic("Error when creating test user " + strconv.Itoa(i) + " (" + utils.TestUsers[i].Email + ")")
		}
	}

	// SLAMS
	if err := repos.slamRepo.CreateWithCreatorTx(ctx, &utils.TestSlams[0], utils.TestUsers[0].Id); err != nil {
		panic("Error when creating test slam 1")
	}
	if err := repos.slamRepo.CreateWithCreatorTx(ctx, &utils.TestSlams[1], utils.TestUsers[1].Id); err != nil {
		panic("Error when creating test slam 2")
	}
	if err := repos.slamRepo.CreateWithCreatorTx(ctx, &utils.TestSlams[2], utils.TestUsers[1].Id); err != nil {
		panic("Error when creating test slam 2")
	}

	// PARTICIPATIONS
	for i := 0; i < len(utils.TestParticipations); i++ {
		if err := repos.partRepo.Create(ctx, &utils.TestParticipations[i]); err != nil {
			panic("Error when creating test participation " + strconv.Itoa(i))
		}
	}

	// STAGES
	for i := 0; i < len(utils.TestStages); i++ {
		if err := repos.stageRepo.Create(ctx, &utils.TestStages[i]); err != nil {
			panic("Error when creating test stage " + strconv.Itoa(i))
		}
	}

	// PERFORMANCES
	for i := 0; i < len(utils.TestPerformances); i++ {
		if err := repos.perfRepo.Create(ctx, &utils.TestPerformances[i]); err != nil {
			panic("Error when creating test performance " + strconv.Itoa(i))
		}
	}

	return r
}
