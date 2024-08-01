package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/validator"
	keeperservice "github.com/kotopesp/sos-kotopes/internal/service/keeper"
	keeperreviewsservice "github.com/kotopesp/sos-kotopes/internal/service/keeper_review"

	v1 "github.com/kotopesp/sos-kotopes/internal/controller/http"
	"github.com/kotopesp/sos-kotopes/internal/core"
	"github.com/kotopesp/sos-kotopes/internal/service/auth"
	"github.com/kotopesp/sos-kotopes/internal/service/name"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kotopesp/sos-kotopes/internal/store/entity"
	keeperstore "github.com/kotopesp/sos-kotopes/internal/store/keeper"
	keeperreviewsstore "github.com/kotopesp/sos-kotopes/internal/store/keeper_review"
	"github.com/kotopesp/sos-kotopes/internal/store/user"

	baseValidator "github.com/go-playground/validator/v10"
	"github.com/kotopesp/sos-kotopes/config"
	"github.com/kotopesp/sos-kotopes/pkg/logger"
	"github.com/kotopesp/sos-kotopes/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	ctx := context.Background()

	// Init logger
	logger.New(cfg.Log.Level)

	// Postgres connection
	pg, err := postgres.New(ctx, cfg.DB.URL)
	if err != nil {
		logger.Log().Fatal(ctx, "error with connection to database: %s", err.Error())
	}
	defer pg.Close(ctx)

	// Stores
	entityStore := entity.New(pg)
	keepersStore := keeperstore.New(pg)
	keeperReviewsStore := keeperreviewsstore.New(pg)
	userStore := user.New(pg)
	// Services
	entityService := name.New(entityStore)
	keeperService := keeperservice.New(keepersStore)
	keeperReviewsService := keeperreviewsservice.New(keeperReviewsStore)
	authService := auth.New(
		userStore,
		core.AuthServiceConfig{
			JWTSecret:      cfg.JWTSecret,
			VKClientID:     cfg.VKClientID,
			VKClientSecret: cfg.VKClientSecret,
			VKCallback:     cfg.VKCallback,
		},
	)

	// Validator
	formValidator := validator.New(ctx, baseValidator.New())

	// HTTP Server
	app := fiber.New(fiber.Config{
		CaseSensitive:            true,
		StrictRouting:            false,
		EnableSplittingOnParsers: true,
	})
	app.Use(recover.New())
	app.Use(cors.New())

	v1.NewRouter(
		app,
		entityService, keeperService, keeperReviewsService,
		formValidator,
		authService,
	)

	logger.Log().Info(ctx, "server was started on %s", cfg.HTTP.Port)
	err = app.ListenTLS(cfg.HTTP.Port, cfg.TLSCert, cfg.TLSKey)
	if err != nil {
		logger.Log().Fatal(ctx, "server was stopped: %s", err.Error())
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Log().Info(ctx, "signal %s received", s.String())
	case <-ctx.Done():
		return
	}

	// Shutdown
	err = app.Shutdown()
	if err != nil {
		logger.Log().Fatal(ctx, "error with shutdown server: %s", err.Error())
	}
}
