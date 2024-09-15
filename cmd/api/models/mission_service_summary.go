package api_models

import (
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionServiceSummaryJson struct {
	Id                string   `json:"id"`
	MissionId         string   `json:"mission_id"`
	Alias             string   `json:"alias"`
	CreatedAt         string   `json:"created_at"`
	AntaresId         string   `json:"antares_id"`
	Description       string   `json:"description"`
	IsImportant       bool     `json:"is_important"`
	NumFireFighters   string   `json:"num_firefighters"`
	NumUnits          string   `json:"num_units"`
	StationName       string   `json:"station_name"`
	NumVehicles       string   `json:"num_vehicles"`
	OperativesAreas   []string `json:"operative_areas"`
	Unharmed          string   `json:"unharmed"`
	Injured           string   `json:"injured"`
	Transported       string   `json:"transported"`
	Deceased          string   `json:"deceased"`
	ServiceDate       string   `json:"service_date"`
	ManualServiceDate string   `json:"manual_service_date"`
}

func ModelToMissionServiceSummaryJson(s models.MissionServiceSummary) *MissionServiceSummaryJson {
	service := MissionServiceSummaryJson{}

	service.MissionId = utils.ParseInt64Sring(s.MissionId)

	service.Id = utils.ParseInt64Sring(s.Id)

	service.Unharmed = utils.ParseInt64Sring(s.Unharmed.Int64)

	service.Injured = utils.ParseInt64Sring(s.Injured.Int64)

	service.Deceased = utils.ParseInt64Sring(s.Deceased.Int64)

	service.Transported = utils.ParseInt64Sring(s.Transported.Int64)

	service.NumUnits = utils.ParseInt64Sring(s.NumUnits.Int64)

	if s.AntaresId.Valid {
		service.AntaresId = utils.ParseInt64Sring(s.AntaresId.Int64)
	}

	if s.Description.Valid {
		service.Description = s.Description.String
	}

	if s.StationName.Valid {
		service.StationName = s.StationName.String
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

	if s.Alias.Valid {
		service.Alias = s.Alias.String
	}

	service.IsImportant = s.IsImportant

	service.OperativesAreas = s.OperativesAreas

	return &service
}

func (s *MissionServiceSummaryJson) ToModel() models.MissionServiceSummary {
	service := models.MissionServiceSummary{}

	service.MissionId = utils.ParseInt64(s.MissionId)

	service.Id = utils.ParseInt64(s.Id)

	service.AntaresId.Int64 = utils.ParseInt64(s.AntaresId)
	service.AntaresId.Valid = true

	service.Description.String = s.Description
	service.Description.Valid = true

	service.StationName.String = s.StationName
	service.StationName.Valid = true

	serviceDate, err := time.Parse("02-01-2006 15:04:05", s.ServiceDate)
	if err == nil {
		service.ServiceDate.Time = serviceDate
		service.ServiceDate.Valid = true
	} else {
		logger.Warn().Err(err).Msg("Problema parseando service date")
	}

	manualServiceDate, err := time.Parse("02-01-2006 15:04:05", s.ManualServiceDate)
	if err == nil {
		service.ManualServiceDate.Time = manualServiceDate
		service.ManualServiceDate.Valid = true
	} else {
		logger.Warn().Err(err).Msg("Problema parseando manual service date")
	}

	service.NumUnits.Int64 = utils.ParseInt64(s.NumUnits)
	service.NumUnits.Valid = true

	service.NumFirefighters.Int64 = utils.ParseInt64(s.NumFireFighters)
	service.NumFirefighters.Valid = true

	service.NumVehicles.Int64 = utils.ParseInt64(s.NumVehicles)
	service.NumVehicles.Valid = true

	service.Unharmed.Int64 = utils.ParseInt64(s.Unharmed)
	service.Unharmed.Valid = true

	service.Injured.Int64 = utils.ParseInt64(s.Injured)
	service.Injured.Valid = true

	service.Transported.Int64 = utils.ParseInt64(s.Transported)
	service.Transported.Valid = true

	service.Deceased.Int64 = utils.ParseInt64(s.Deceased)
	service.Deceased.Valid = true

	service.IsImportant = s.IsImportant

	service.OperativesAreas = s.OperativesAreas

	return service

}
