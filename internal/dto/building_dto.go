package dto


type BuildingDTO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Year   int    `json:"year"`
	Floors int    `json:"floors"`
}

type FiltersBuildingDTO struct{
	City   *string `json:"city"`
	Year   *int    `json:"year"`
	Floors *int    `json:"floors"`
}

type CreateBuildingDTO struct {
	Name   string `json:"name"`
	City   string `json:"city"`
	Year   int    `json:"year"`
	Floors int    `json:"floors"`
}