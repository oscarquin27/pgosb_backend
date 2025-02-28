package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionReportSummaryJson struct {
	MissionId           string `json:"mission_id"`
	AntaresId           string `json:"antares_id"`
	AntaresName         string `json:"antares_name"`
	AntaresType         string `json:"antares_type"`
	StationId           string `json:"station_id"`
	StationName         string `json:"station_name"`
	StationAbbreviation string `json:"station_abbreviation"`
	Count               int64  `json:"count,omitempty"`
}

func ModelToMissionReportSummaryJson(s models.MissionReportSummary) *MissionReportSummaryJson {
	report := MissionReportSummaryJson{}

	report.MissionId = utils.ParseInt64String(s.MissionId)

	if s.AntaresId.Valid {
		report.AntaresId = utils.ParseInt64String(s.AntaresId.Int64)
	}

	if s.AntaresName.Valid {
		report.AntaresName = s.AntaresName.String
	}

	if s.AntaresType.Valid {
		report.AntaresType = s.AntaresType.String
	}

	if s.StationId.Valid {
		report.StationId = utils.ParseInt64String(s.StationId.Int64)
	}

	if s.StationName.Valid {
		report.StationName = s.StationName.String
	}

	if s.StationAbbreviation.Valid {
		report.StationAbbreviation = s.StationAbbreviation.String
	}

	return &report
}

func (s *MissionReportSummaryJson) ToModel() models.MissionReportSummary {
	report := models.MissionReportSummary{}

	report.MissionId = utils.ParseInt64(s.MissionId)

	if s.AntaresId != "" {
		report.AntaresId.Int64 = utils.ParseInt64(s.AntaresId)
		report.AntaresId.Valid = true
	}

	if s.AntaresName != "" {
		report.AntaresName.String = s.AntaresName
		report.AntaresName.Valid = true
	}

	if s.AntaresType != "" {
		report.AntaresType.String = s.AntaresType
		report.AntaresType.Valid = true
	}

	if s.StationId != "" {
		report.StationId.Int64 = utils.ParseInt64(s.StationId)
		report.StationId.Valid = true
	}

	if s.StationName != "" {
		report.StationName.String = s.StationName
		report.StationName.Valid = true
	}

	if s.StationAbbreviation != "" {
		report.StationAbbreviation.String = s.StationAbbreviation
		report.StationAbbreviation.Valid = true
	}

	return report
}
