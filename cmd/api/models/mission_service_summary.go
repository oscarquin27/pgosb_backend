package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionServiceSummaryJson struct {
	Id                string `json:"id"`
	Alias             string `json:"alias"`
	CreatedAt         string `json:"created_at"`
	ServiceId         string `json:"service_id"`
	AntaresId         string `json:"antares_id"`
	Description       string `json:"description"`
	ServiceDate       string `json:"service_date"`
	ManualServiceDate string `json:"manual_service_date"`
	NumFireFighters   string `json:"num_firefighters"`
	NumVehicles       string `json:"num_vehicles"`
	StationName       string `json:"station_name"`
}

func ModelToMissionServiceSummaryJson(s models.MissionServiceSummary) *MissionServiceSummaryJson {
	service := MissionServiceSummaryJson{}

	service.Id = utils.ParseInt64Sring(s.Id)

	if s.ServiceId.Valid {
		service.ServiceId = utils.ParseInt64Sring(s.ServiceId.Int64)
	}

	if s.AntaresId.Valid {
		service.AntaresId = utils.ParseInt64Sring(s.AntaresId.Int64)
	}

	if s.Alias.Valid {
		service.Alias = s.Alias.String
	}

	if s.Description.Valid {
		service.Description = s.Description.String
	}

	if s.StationName.Valid {
		service.StationName = s.StationName.String
	}

	if s.CreatedAt.Valid {
		service.CreatedAt = s.CreatedAt.Time.Format("02-01-2006 15:04:05")
	}

	if s.ServiceDate.Valid {
		service.ServiceDate = s.ServiceDate.Time.Format("02-01-2006 15:04:05")
	}

	if s.ManualServiceDate.Valid {
		service.ManualServiceDate = s.ManualServiceDate.Time.Format("02-01-2006 15:04:05")
	}

	if s.NumFirefighters.Valid {
		service.NumFireFighters = utils.ParseInt64Sring(s.NumFirefighters.Int64)
	}

	if s.NumVehicles.Valid {
		service.NumVehicles = utils.ParseInt64Sring(s.NumVehicles.Int64)
	}

	return &service
}

func (s *MissionServiceSummaryJson) ToModel() models.MissionServiceSummary {
	service := models.MissionServiceSummary{}

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
