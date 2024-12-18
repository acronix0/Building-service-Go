package service

import (
	"context"

	"github.com/acronix0/Building-service-Go/internal/dto"
)

type ServiceManager interface {
	Building() Building
}

type Building interface {
	Create(ctx context.Context, building dto.CreateBuildingDTO) error
	Get(ctx context.Context, filters dto.FiltersBuildingDTO) ([]dto.BuildingDTO, error)
}
