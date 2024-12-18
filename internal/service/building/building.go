package building

import (
	"github.com/acronix0/Building-service-Go/internal/repository"
)

type BuildingService struct {
	repository repository.Building
}

func NewBuildingService(repository repository.Building) *BuildingService {
	return &BuildingService{repository: repository}
}


