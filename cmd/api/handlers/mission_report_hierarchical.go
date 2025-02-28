package mission_report_hierarchical_handler

import (
	"fdms/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionIds struct {
	MissionIds []int64 `json:"mission_ids"`
}

// MissionReportHierarchicalController handles hierarchical mission report requests
type MissionReportHierarchicalController struct {
	service services.MissionReportHierarchicalService
}

// NewMissionReportHierarchicalController creates a new hierarchical report controller
func NewMissionReportHierarchicalController(service services.MissionReportHierarchicalService) *MissionReportHierarchicalController {
	return &MissionReportHierarchicalController{
		service: service,
	}
}

// RegisterRoutes registers the routes for the hierarchical report controller
func (c *MissionReportHierarchicalController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/mission-reports")
	{
		group.POST("/hierarchical/in", c.GetHierarchicalReport)
	}
}

// GetHierarchicalReport handles the request for a hierarchical mission report
func (c *MissionReportHierarchicalController) GetHierarchicalReport(ctx *gin.Context) {
	// Get optional date range filters from query parameters
	// startDate := ctx.Query("start_date")
	// endDate := ctx.Query("end_date")

	var missionId []string

	if err := ctx.ShouldBindJSON(&missionId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(missionId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mission_ids is required"})
		return
	}

	report, err := c.service.GetHierarchicalReport(missionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}
