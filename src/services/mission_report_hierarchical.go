package services

import "fdms/src/models"

// MissionReportHierarchicalService defines the interface for hierarchical mission reports
type MissionReportHierarchicalService interface {
	// GetHierarchicalReport generates a hierarchical report with regions, stations, and missions
	// startDate and endDate are optional filters (can be empty strings)
	GetHierarchicalReport(missionIds []string) (*models.HierarchicalReport, error)
}
