package building

import (
	"context"
	"fmt"

	"github.com/acronix0/Building-service-Go/internal/repository/building/model"
)


func (r *BuildingRepo) Get(ctx context.Context, filters model.BuildingFilters) ([]model.Building, error) {
	query := `
		SELECT id, name, city, year, floors
		FROM buildings
		WHERE 1=1
	`
	args := []interface{}{}
	argIdx := 1

	if filters.City != nil {
		query += fmt.Sprintf(" AND city = $%d", argIdx)
		args = append(args, *filters.City)
		argIdx++
	}
	if filters.Year != nil {
		query += fmt.Sprintf(" AND year = $%d", argIdx)
		args = append(args, *filters.Year)
		argIdx++
	}
	if filters.Floors != nil {
		query += fmt.Sprintf(" AND floors = $%d", argIdx)
		args = append(args, *filters.Floors)
		argIdx++
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []model.Building
	for rows.Next() {
		var b model.Building
		if err := rows.Scan(&b.ID, &b.Name, &b.City, &b.Year, &b.Floors); err != nil {
			return nil, err
		}
		buildings = append(buildings, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return buildings, nil
}