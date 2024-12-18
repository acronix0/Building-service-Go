package building

import (
	"context"

	"github.com/acronix0/Building-service-Go/internal/dto"
	"github.com/acronix0/Building-service-Go/internal/mapper"
)

func (s *BuildingService) Get(ctx context.Context, filters dto.FiltersBuildingDTO) ([]dto.BuildingDTO, error) {
	buildings, err := s.repository.Get(ctx, mapper.MapFiltersDTOToModel(filters))
	if err != nil {
		return []dto.BuildingDTO{}, err
	}

	result := make([]dto.BuildingDTO, len(buildings))
	for _, building := range buildings {
		result = append(result, mapper.MapModelToDTO(building))
	}

	return result, nil
}