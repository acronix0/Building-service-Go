package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/acronix0/Building-service-Go/internal/config"
	"github.com/acronix0/Building-service-Go/internal/database"
	delivery "github.com/acronix0/Building-service-Go/internal/delivery/http"
	"github.com/acronix0/Building-service-Go/internal/server"
)

type App struct {
	serviceProvider *serviceProvider
	logger          *slog.Logger
	config          *config.Config
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// @title Building API
// @version 1.0
// @description API for managing buildings.

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func (a *App) Run() {
	handlers := delivery.NewHandler(a.serviceProvider.ServiceManager(), a.logger)
	srv := server.NewServer(a.config, handlers.Init(a.config))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("error occurred while running http server: %s\n", err.Error())
		}
	}()

	a.logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		a.logger.Error("failed to stop server: %v", err)
	}
}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		//a.InitMigrations,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.MustLoad()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	var log *slog.Logger
	switch a.config.AppEnv {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	a.logger = log
	return nil
}
func (a *App) InitMigrations(_ context.Context) error {
	return database.InitMigrations(
		a.config.MigrationsPath,
		a.config.DatabaseConnection.Host,
		a.config.DatabaseConnection.User,
		a.config.DatabaseConnection.Name,
		a.config.DatabaseConnection.Password,
		a.config.DatabaseConnection.Port,
	)
}
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.config, a.logger)
	return nil
}
