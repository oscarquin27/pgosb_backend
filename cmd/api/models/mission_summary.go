package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionSummaryJson struct {
	Id              string `json:"id"`
	Alias           string `json:"alias"`
	CreatedAt       string `json:"created_at"`
	NumServices     string `json:"num_services"`
	NumFireFighters string `json:"num_firefighters"`
	NumVehicles     string `json:"num_vehicles"`
}

func ModelToMissionSummaryJson(s models.MissionSummary) *MissionSummaryJson {
	service := MissionSummaryJson{}

	service.Id = utils.ParseInt64Sring(s.Id)

	if s.Alias.Valid {
		service.Alias = s.Alias.String
	}

	if s.CreatedAt.Valid {
		service.CreatedAt = s.CreatedAt.Time.Format("02-01-2006 15:04:05")
	}

	if s.NumServices.Valid {
		service.NumServices = utils.ParseInt64Sring(s.NumServices.Int64)
	}

	if s.NumVehicles.Valid {
		service.NumVehicles = utils.ParseInt64Sring(s.NumVehicles.Int64)
	}

	if s.NumFireFighters.Valid {
		service.NumFireFighters = utils.ParseInt64Sring(s.NumFireFighters.Int64)
	}

	return &service
}

func (s *MissionSummaryJson) ToModel() models.MissionSummary {
	service := models.MissionSummary{}

	// createdAt, err := time.Parse("02-01-2006 15:04:05", s.CreatedAt)

	// if err == nil {
	// 	service.CreatedAt.Time = createdAt
	// 	service.CreatedAt.Valid = true
	// } else {
	// 	logger.Warn().Err(err).Msg("Problema parseando created at date")
	// }

	service.Alias.String = s.Alias
	service.Alias.Valid = true

	return service
}
