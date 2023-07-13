//go:build wireinject
// +build wireinject

package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/internal/api"
	"go-clean-architecture-example/internal/app"
	"go-clean-architecture-example/internal/common/errors"
	"go-clean-architecture-example/internal/common/logger"
	"go-clean-architecture-example/internal/infrastructure/notification"
	"go-clean-architecture-example/internal/infrastructure/persistence"
	loggerPkg "go-clean-architecture-example/pkg/logger"

	"go-clean-architecture-example/internal/router"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"go-clean-architecture-example/docs"
	"go-clean-architecture-example/internal/probes"
	"os"
	"time"
)

// Server struct
type Server struct {
	app    *fiber.App
	cfg    *config.Configuration
	logger loggerPkg.Logger
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		router.Set,
		api.Set,
		app.Set,
		persistence.Set,
		notification.Set,
		probes.Set,
		logger.Set,
	)))
}

// @title  My SERVER
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email minkj1992@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
func NewServer(
	cfg *config.Configuration,
	cragRouter router.CragRouter,
	healthCheckApp probes.HealthCheckApplication,
	logger loggerPkg.Logger) *Server {

	app := fiber.New(fiber.Config{
		ErrorHandler: errors.CustomErrorHandler,
		ReadTimeout:  time.Second * cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * cfg.Server.WriteTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(fiberlog.New(fiberlog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	setSwagger(cfg.Server.BaseURI)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// health check endpoint
	app.Get("/liveness", func(c *fiber.Ctx) error {
		result := healthCheckApp.LiveEndpoint()
		if result.Status {
			return c.Status(fiber.StatusOK).JSON(result)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	})

	app.Get("/readiness", func(c *fiber.Ctx) error {
		result := healthCheckApp.ReadyEndpoint()
		if result.Status {
			return c.Status(fiber.StatusOK).JSON(result)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	cragRouter.Init(&v1)

	return &Server{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Configuration {
	return serv.cfg
}

func (serv Server) Logger() loggerPkg.Logger {
	return serv.logger
}

func setSwagger(baseURI string) {
	docs.SwaggerInfo.Title = "Go Clean Architecture Example ✈️"
	docs.SwaggerInfo.Description = "This is a go clean architecture example."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = baseURI
	docs.SwaggerInfo.BasePath = "/api/v1"
}
