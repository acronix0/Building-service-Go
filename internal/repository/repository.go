package repository

import (
	"context"
	"database/sql"

	"github.com/acronix0/Building-service-Go/internal/repository/building/model"
	"github.com/acronix0/Building-service-Go/internal/repository/building"
)

type Building interface {
	Create(ctx context.Context, building model.Building) error
	Get(ctx context.Context, filters model.BuildingFilters) ([]model.Building, error)
}

type RepositoryManager interface {
	Building() Building
}

var _ Building = (*building.BuildingRepo)(nil)

type repositories struct {
	db     *sql.DB
	building   Building
}

func NewRepositoryManager(db *sql.DB) *repositories {
	return &repositories{db: db}
}
func (r *repositories) DB() *sql.DB {
	return r.db
}
func (r *repositories) Building() Building {
	if r.building == nil {
		r.building = building.NewBuildingitory(r.DB())
	}
	return r.building
}

