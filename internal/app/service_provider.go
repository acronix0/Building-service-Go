package app

import (
	"log/slog"
	"os"

	"github.com/acronix0/Building-service-Go/internal/config"
	"github.com/acronix0/Building-service-Go/internal/database"
	"github.com/acronix0/Building-service-Go/internal/repository"
	"github.com/acronix0/Building-service-Go/internal/service"
	"github.com/acronix0/Building-service-Go/internal/service/building"
)

type serviceProvider struct {
	config            *config.Config
	dataBase          *database.Database
	logger            *slog.Logger
	serviceManager    service.ServiceManager
	repositoryManager repository.RepositoryManager
}

func newServiceProvider(cfg *config.Config, logger *slog.Logger) *serviceProvider {
	return &serviceProvider{config: cfg, logger: logger}
}

func (s *serviceProvider) Config() *config.Config {
	return s.config
}
func (s *serviceProvider) Database() *database.Database {
	if s.dataBase == nil {
		cfg := s.Config()
		db, err := database.NewDatabase(
			cfg.DatabaseConnection.Port,
			cfg.DatabaseConnection.Host,
			cfg.DatabaseConnection.User,
			cfg.DatabaseConnection.Password,
			cfg.DatabaseConnection.Name)
		if err != nil {
			panic(err)
		}
		s.dataBase = db
	}

	return s.dataBase
}
func (s *serviceProvider) Logger() *slog.Logger {
	if s.logger == nil {
		s.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return s.logger
}
func (s *serviceProvider) ServiceManager() service.ServiceManager {
	if s.serviceManager == nil {
		s.serviceManager = building.NewService(s.RepositoryManager())

	}

	return s.serviceManager
}

func (s *serviceProvider) RepositoryManager() repository.RepositoryManager {
	if s.repositoryManager == nil {
		s.repositoryManager = repository.NewRepositoryManager(s.Database().GetDB())
	}
	return s.repositoryManager
}


