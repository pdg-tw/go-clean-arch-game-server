package app

import (
	"go-clean-arch-game-server/internal/app/crag/commands"
	"go-clean-arch-game-server/internal/app/crag/queries"
	memcommands "go-clean-arch-game-server/internal/app/member/commands"
	"go-clean-arch-game-server/internal/common/metrics"
	"go-clean-arch-game-server/internal/domain/entities/crag"
	"go-clean-arch-game-server/internal/domain/entities/member"
	"go-clean-arch-game-server/internal/domain/entities/notification"
	"go-clean-arch-game-server/pkg/logger"
	"go-clean-arch-game-server/pkg/time"
	"go-clean-arch-game-server/pkg/uuid"
)

// Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler queries.GetAllCragsRequestHandler
	GetCragHandler     queries.GetCragRequestHandler
}

// Commands Contains all available command handlers of this app
type Commands struct {
	AddCragHandler    commands.AddCragRequestHandler
	UpdateCragHandler commands.UpdateCragRequestHandler
	DeleteCragHandler commands.DeleteCragRequestHandler
	AddMemberHandler  memcommands.AddMemberRequestHandler
}

type Application struct {
	Queries  Queries
	Commands Commands
}

func NewApplication(cragRepo crag.Repository, memberRepo member.Repository, ns notification.Service, logger logger.Logger) Application {
	// init base
	metricsClient := metrics.NoOp{}
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	return Application{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragRepo, logger, metricsClient),
			GetCragHandler:     queries.NewGetCragRequestHandler(cragRepo, logger, metricsClient),
		},
		Commands: Commands{
			AddCragHandler:    commands.NewAddCragRequestHandler(up, tp, cragRepo, ns, logger, metricsClient),
			UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragRepo, logger, metricsClient),
			DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragRepo, logger, metricsClient),
			AddMemberHandler:  memcommands.NewAddMemberRequestHandler(up, tp, memberRepo, ns, logger, metricsClient),
		},
	}
}
