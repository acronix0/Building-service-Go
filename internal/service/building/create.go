package building

import (
	"context"

	"github.com/acronix0/Building-service-Go/internal/dto"
	"github.com/acronix0/Building-service-Go/internal/mapper"
)

func (s *BuildingService) Create(ctx context.Context, building dto.CreateBuildingDTO) error {
	err := s.repository.Create(ctx, mapper.MapCreateDTOToModel(building))
	if err != nil {
		return err
	}

	return nil
}