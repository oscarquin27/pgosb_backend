package services

import "fdms/src/models"

type MissionReportAggregationsService interface {
	// GetAntaresStationByMissionIds returns aggregated data grouped by antares and station
	GetAntaresStationByMissionIds(missionIds []string) ([]models.AntaresStationAggregation, error)

	// GetAntaresByMissionIds returns aggregated data grouped by antares
	GetAntaresByMissionIds(missionIds []string) ([]models.AntaresAggregation, error)

	// GetStationByMissionIds returns aggregated data grouped by station
	GetStationByMissionIds(missionIds []string) ([]models.StationAggregation, error)

	// GetAntaresTypeByMissionIds returns aggregated data grouped by antares type
	GetAntaresTypeByMissionIds(missionIds []string) ([]models.AntaresTypeAggregation, error)
}
