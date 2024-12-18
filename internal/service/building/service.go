package building

import (
	"github.com/acronix0/Building-service-Go/internal/repository"
	"github.com/acronix0/Building-service-Go/internal/service"
)

type services struct {
	BuildingService service.Building
}

func NewService(repoManager repository.RepositoryManager) *services {
	return &services{
		BuildingService: NewBuildingService(repoManager.Building()),
	}
}


func (s *services) Building() service.Building {
	return s.BuildingService
}
