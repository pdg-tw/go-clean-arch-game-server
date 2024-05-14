package queries

import (
	"context"
	"go-clean-arch-game-server/internal/common/decorator"
	"go-clean-arch-game-server/internal/common/utils"
	dto "go-clean-arch-game-server/internal/domain/dto/crag"
	"go-clean-arch-game-server/internal/domain/entities/crag"
	"go-clean-arch-game-server/pkg/logger"
)

type GetCragRequestHandler decorator.QueryHandler[*dto.GetCragRequest, *dto.GetCragResult]

type getCragRequestHandler struct {
	repo crag.Repository
}

// NewGetCragRequestHandler Handler Constructor
func NewGetCragRequestHandler(
	repo crag.Repository,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) GetCragRequestHandler {
	return decorator.ApplyQueryDecorators[*dto.GetCragRequest, *dto.GetCragResult](
		getCragRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handlers the GetCragRequest query
func (h getCragRequestHandler) Handle(ctx context.Context, query *dto.GetCragRequest) (*dto.GetCragResult, error) {
	var result dto.GetCragResult

	cragData, err := h.repo.GetByID(query.CragID)
	if err != nil {
		return &result, err
	}
	err = utils.BindingStruct(cragData, &result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}
