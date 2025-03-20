package mission_report_agregations_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MissionReportAggregationsController handles aggregation mission report requests
type MissionReportAggregationsController struct {
	service services.MissionReportAggregationsService
}

// NewMissionReportAggregationsController creates a new aggregations report controller
func NewMissionReportAggregationsController(service services.MissionReportAggregationsService) *MissionReportAggregationsController {
	return &MissionReportAggregationsController{
		service: service,
	}
}

// RegisterRoutes registers the routes for the aggregations report controller
func (c *MissionReportAggregationsController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/mission-reports/aggregations")
	{
		group.POST("/antares-station/in", c.GetAntaresStationAggregation)
		group.POST("/antares/in", c.GetAntaresAggregation)
		group.POST("/station/in", c.GetStationAggregation)
		group.POST("/antares-type/in", c.GetAntaresTypeAggregation)
		group.POST("/municipality/in", c.GetMunicipalityAggregation)
		group.POST("/parish/in", c.GetParishAggregation)
	}
}

// GetAntaresStationAggregation handles the request for antares-station aggregation
func (c *MissionReportAggregationsController) GetAntaresStationAggregation(ctx *gin.Context) {
	var missionIds []string

	if err := ctx.ShouldBindJSON(&missionIds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionIds) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetAntaresStationByMissionIds(missionIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to JSON format
	jsonReport := api_models.ModelToAntaresStationAggregationJsonList(report)
	ctx.JSON(http.StatusOK, jsonReport)
}

// GetAntaresAggregation handles the request for antares aggregation
func (c *MissionReportAggregationsController) GetAntaresAggregation(ctx *gin.Context) {
	var missionIds []string

	if err := ctx.ShouldBindJSON(&missionIds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionIds) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetAntaresByMissionIds(missionIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to JSON format
	jsonReport := api_models.ModelToAntaresAggregationJsonList(report)
	ctx.JSON(http.StatusOK, jsonReport)
}

func (c *MissionReportAggregationsController) GetMunicipalityAggregation(ctx *gin.Context) {
	var missionIds []string

	if err := ctx.ShouldBindJSON(&missionIds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionIds) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetMunicipalityOriginByMissionIds(missionIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to JSON format
	jsonReport := api_models.ModelToMunicipalityTypeAggregationJsonList(report)
	ctx.JSON(http.StatusOK, jsonReport)
}

func (c *MissionReportAggregationsController) GetParishAggregation(ctx *gin.Context) {
	var missionIds []string

	if err := ctx.ShouldBindJSON(&missionIds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionIds) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetParishOriginByMissionIds(missionIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to JSON format
	jsonReport := api_models.ModelToParishTypeAggregationJsonList(report)
	ctx.JSON(http.StatusOK, jsonReport)
}

// GetStationAggregation handles the request for station aggregation
func (c *MissionReportAggregationsController) GetStationAggregation(ctx *gin.Context) {
	var missionIds []string

	if err := ctx.ShouldBindJSON(&missionIds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionIds) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetStationByMissionIds(missionIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to JSON format
	jsonReport := api_models.ModelToStationAggregationJsonList(report)
	ctx.JSON(http.StatusOK, jsonReport)
}

// GetAntaresTypeAggregation handles the request for antares-type aggregation
func (c *MissionReportAggregationsController) GetAntaresTypeAggregation(ctx *gin.Context) {
	var missionIds []string

	if err := ctx.ShouldBindJSON(&missionIds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionIds) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetAntaresTypeByMissionIds(missionIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to JSON format
	jsonReport := api_models.ModelToAntaresTypeAggregationJsonList(report)
	ctx.JSON(http.StatusOK, jsonReport)
}
