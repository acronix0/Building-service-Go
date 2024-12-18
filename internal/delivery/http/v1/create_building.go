package v1

import (
	"log/slog"
	"net/http"

	"github.com/acronix0/Building-service-Go/internal/dto"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// createBuilding handles the POST request for adding a new building.
// @Summary Create Building
// @Tags Buildings
// @Description Add a new building to the database.
// @Accept json
// @Produce json
// @Param input body dto.BuildingDTO true "Building details"
// @Success 201 {object} dto.BuildingDTO "Successfully created building"
// @Failure 400 {object} Response "Invalid input data"
// @Failure 500 {object} Response "Failed to add building"
// @Router /building [post]
func (h *Handler) createBuilding(c *gin.Context) {
	const op = "handler.v1.createBuilding"
	logger := h.logger.With(
		slog.String("op", op),
	)

	var buildingDTO dto.BuildingDTO
	if err := c.ShouldBindJSON(&buildingDTO); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input data")
		logger.Error("Invalid input data", slog.String("error", err.Error()))
		return
	}

	err := h.services.Building().Create(
		c.Request.Context(), 
		buildingDTO)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to add building")
		logger.Error("Failed to add building", slog.String("error", err.Error()))
		return
	}

	logger.Debug("Building added successfully")
	c.JSON(http.StatusCreated, buildingDTO)
}
