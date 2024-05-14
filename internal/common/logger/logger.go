package logger

import (
	"github.com/google/wire"
	"go-clean-arch-game-server/config"
	"go-clean-arch-game-server/pkg/logger"
)

var Set = wire.NewSet(
	NewLoggerAplication,
)

// NewHandler Constructor
func NewLoggerAplication(cfg *config.Configuration) logger.Logger {
	return logger.NewApiLogger(cfg)
}
