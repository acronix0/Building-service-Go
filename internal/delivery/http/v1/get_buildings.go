package v1

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/acronix0/Building-service-Go/internal/dto"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// getBuildings handles the GET request for fetching a list of buildings.
// @Summary Get Buildings
// @Tags Buildings
// @Description Get a filtered list of buildings.
// @Produce json
// @Param city query string false "Filter by city"
// @Param year query int false "Filter by year"
// @Param floors query int false "Filter by number of floors"
// @Success 200 {array} dto.BuildingDTO "List of buildings"
// @Failure 400 {object} Response "Invalid query parameters"
// @Failure 500 {object} Response "Internal server error"
// @Router /building [get]
func (h *Handler) getBuildings(c *gin.Context) {
	const op = "handler.v1.getBuildings"
	logger := h.logger.With(
		slog.String("op", op),
	)

	logger.Info("Start handling getBuildings request")

	filters := dto.FiltersBuildingDTO{
		City:   parseStringQuery(c, "city"),
		Year:   parseIntQuery(c, "year"),
		Floors: parseIntQuery(c, "floors"),
	}


	buildings, err := h.services.Building().Get(c.Request.Context(), filters)
	if err != nil {
		logger.Error("Failed to get buildings", slog.String("error", err.Error()))
		newResponse(c, http.StatusInternalServerError, "Failed to get buildings")
		return
	}

	if len(buildings) == 0 {
		logger.Info("No buildings found")
		c.JSON(http.StatusOK, []dto.BuildingDTO{})
		return
	}

	logger.Info("Successfully retrieved buildings", slog.Int("count", len(buildings)))
	c.JSON(http.StatusOK, buildings)
}

func parseIntQuery(c *gin.Context, key string) *int {
    if value := c.Query(key); value != "" {
        if parsedValue, err := strconv.Atoi(value); err == nil {
            return &parsedValue
        }
    }
    return nil
}

func parseStringQuery(c *gin.Context, key string) *string {
    if value := c.Query(key); value != "" {
        return &value
    }
    return nil
}