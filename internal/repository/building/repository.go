package building

import (
	"database/sql"
)

type BuildingRepo struct {
	db *sql.DB
}

func NewBuildingitory(db *sql.DB) *BuildingRepo {
	return &BuildingRepo{db: db}
}
