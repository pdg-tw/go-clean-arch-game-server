package memcommands

import (
	"context"
	"go-clean-arch-game-server/internal/common/decorator"
	dto "go-clean-arch-game-server/internal/domain/dto/member"
	"go-clean-arch-game-server/internal/domain/entities/member"
	"go-clean-arch-game-server/internal/domain/entities/notification"
	"go-clean-arch-game-server/pkg/logger"
	timePkg "go-clean-arch-game-server/pkg/time"
	uuidPkg "go-clean-arch-game-server/pkg/uuid"
)

type AddMemberRequestHandler decorator.CommandHandler[*dto.AddMemberRequest]

type addMemberRequestHandler struct {
	uuidProvider        uuidPkg.Provider
	timeProvider        timePkg.Provider
	repo                member.Repository
	notificationService notification.Service
}

// NewAddCragRequestHandler Initializes an AddCommandHandler
func NewAddMemberRequestHandler(
	uuidProvider uuidPkg.Provider,
	timeProvider timePkg.Provider,
	repo member.Repository,
	notificationService notification.Service,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) AddMemberRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.AddMemberRequest](
		addMemberRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService},
		logger,
		metricsClient,
	)
}

// Handle Handles the AddCragRequest
func (h addMemberRequestHandler) Handle(ctx context.Context, req *dto.AddMemberRequest) error {
	c := member.Member{
		ID:        h.uuidProvider.NewUUID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: h.timeProvider.Now(),
	}
	err := h.repo.Add(c)
	if err != nil {
		return err
	}
	n := notification.Notification{
		Subject: "New Member added",
		Message: "A new Member with name '" + c.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
