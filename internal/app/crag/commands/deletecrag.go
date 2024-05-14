package commands

import (
	"context"
	"fmt"
	"go-clean-arch-game-server/internal/common/decorator"
	dto "go-clean-arch-game-server/internal/domain/dto/crag"
	"go-clean-arch-game-server/internal/domain/entities/crag"
	"go-clean-arch-game-server/pkg/logger"
)

// DeleteCragRequestHandler Handler Struct with Dependencies
type DeleteCragRequestHandler decorator.CommandHandler[*dto.DeleteCragRequest]

type deleteCragRequestHandler struct {
	repo crag.Repository
}

// NewDeleteCragRequestHandler Handler constructor
func NewDeleteCragRequestHandler(
	repo crag.Repository,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) DeleteCragRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.DeleteCragRequest](
		deleteCragRequestHandler{repo: repo},
		logger,
		metricsClient)
}

// Handle Handlers the DeleteCragRequest request
func (h deleteCragRequestHandler) Handle(ctx context.Context, command *dto.DeleteCragRequest) error {
	crag, err := h.repo.GetByID(command.CragID)
	if crag == nil {
		return fmt.Errorf("the provided crag id does not exist")
	}
	if err != nil {
		return err
	}
	return h.repo.Delete(command.CragID)
}
