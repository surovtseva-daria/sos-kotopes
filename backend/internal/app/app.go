package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/validator"
	rolesService "github.com/kotopesp/sos-kotopes/internal/service/role"
	usersService "github.com/kotopesp/sos-kotopes/internal/service/user"
	rolesStore "github.com/kotopesp/sos-kotopes/internal/store/role"
	userFavouriteStore "github.com/kotopesp/sos-kotopes/internal/store/userfavourite"

	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/kotopesp/sos-kotopes/config"
	v1 "github.com/kotopesp/sos-kotopes/internal/controller/http"
	"github.com/kotopesp/sos-kotopes/internal/core"
	"github.com/kotopesp/sos-kotopes/internal/service/auth"
	"github.com/kotopesp/sos-kotopes/internal/store/user"
	"github.com/kotopesp/sos-kotopes/pkg/logger"
	"github.com/kotopesp/sos-kotopes/pkg/postgres"

	commentservice "github.com/kotopesp/sos-kotopes/internal/service/comment_service"
	postservice "github.com/kotopesp/sos-kotopes/internal/service/post"
	animalstore "github.com/kotopesp/sos-kotopes/internal/store/animal"
	commentstore "github.com/kotopesp/sos-kotopes/internal/store/comment_store"
	poststore "github.com/kotopesp/sos-kotopes/internal/store/post"
	postfavouritestore "github.com/kotopesp/sos-kotopes/internal/store/postfavourite"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	psql "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	DatabaseUrl      = "postgres://postgres:postgres@postgres:5432/soskot?sslmode=disable"
	Db               = "postgres"
	PathToMogrations = "file:///app/internal/data/"
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

	// Migrate up
	db, err := sql.Open(Db, DatabaseUrl)
	if err != nil {
		logger.Log().Fatal(ctx, err.Error())
	}
	driver, err := psql.WithInstance(db, &psql.Config{})
	if err != nil {
		logger.Log().Fatal(ctx, err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		PathToMogrations,
		Db, driver)
	if err != nil {
		logger.Log().Fatal(ctx, err.Error())
	}
	if m.Up() != nil {
		logger.Log().Fatal(ctx, "error with up migrations for database")
	}

	// Stores
	userStore := user.New(pg)
	commentStore := commentstore.New(pg)
	roleStore := rolesStore.New(pg)
	favouriteUserStore := userFavouriteStore.New(pg)
	postStore := poststore.New(pg)
	postFavouriteStore := postfavouritestore.New(pg)
	animalStore := animalstore.New(pg)

	// Services
	commentService := commentservice.New(
		commentStore,
		postStore,
	)
	roleService := rolesService.New(roleStore, userStore)
	userService := usersService.New(userStore, favouriteUserStore)
	authService := auth.New(
		userStore,
		core.AuthServiceConfig{
			JWTSecret:            cfg.JWTSecret,
			VKClientID:           cfg.VKClientID,
			VKClientSecret:       cfg.VKClientSecret,
			VKCallback:           cfg.VKCallback,
			AccessTokenLifetime:  cfg.AccessTokenLifetime,
			RefreshTokenLifetime: cfg.RefreshTokenLifetime,
		},
	)
	postService := postservice.New(postStore, postFavouriteStore, animalStore, userStore)

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
		authService,
		commentService,
		postService,
		userService,
		roleService,
		formValidator,
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
