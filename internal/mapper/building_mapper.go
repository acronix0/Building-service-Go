package mapper

import (
	"github.com/acronix0/Building-service-Go/internal/dto"
	"github.com/acronix0/Building-service-Go/internal/repository/building/model"
)

func MapDTOToModel(dto dto.BuildingDTO) model.Building {
	return model.Building{
		Name:   dto.Name,
		City:   dto.City,
		Year:   dto.Year,
		Floors: dto.Floors,
	}
}

func MapModelToDTO(building model.Building) dto.BuildingDTO {
	return dto.BuildingDTO{
		ID:     building.ID,
		Name:   building.Name,
		City:   building.City,
		Year:   building.Year,
		Floors: building.Floors,
	}
}

func MapFiltersDTOToModel(filters dto.FiltersBuildingDTO) model.BuildingFilters {
	return model.BuildingFilters{
		City:   filters.City,
		Year:   filters.Year,
		Floors: filters.Floors,
	}
}

func MapFiltersModelToDTO(filters model.BuildingFilters) dto.FiltersBuildingDTO {
	return dto.FiltersBuildingDTO{
		City:   filters.City,
		Year:   filters.Year,
		Floors: filters.Floors,
	}
}
