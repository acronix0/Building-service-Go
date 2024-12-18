package building

import (
	"context"
	
	"github.com/acronix0/Building-service-Go/internal/repository/building/model"

)


func (r *BuildingRepo) Create(ctx context.Context, building model.Building) error {
	query := `
		INSERT INTO buildings (name, city, year, floors)
		VALUES ($1, $2, $3, $4) RETURNING id
	`

	err := r.db.QueryRow(
		query, 
		building.Name, 
		building.City, 
		building.Year, 
		building.Floors,
		).Scan(&building.ID)

	if err != nil {
		return err
	}

	return nil
}
