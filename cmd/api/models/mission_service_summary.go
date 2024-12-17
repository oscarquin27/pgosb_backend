package api_models

import (
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionServiceSummaryJson struct {
	Id                string `json:"id"`
	Alias 			  string `json:"alias"`
	Code              string `json:"code"`
	CreatedAt         string   `json:"created_at"`
	NumServices       string   `json:"num_services"`
	Unharmed          string   `json:"unharmed"`
	Injured           string   `json:"injured"`
	Transported       string   `json:"transported"`
	Deceased          string   `json:"deceased"`
	NumVehicles       string   `json:"num_vehicles"`
	NumFirefighters   string   `json:"num_firefighters"`
	OperativesAreas   []string       `json:"operative_areas"`
	NumAuthorities	  string   `json:"num_authorities"`
	NumAuthorityServices  string   `json:"num_authority_services"`
	NumAuthorityPerson  string   `json:"num_authority_person"`
	NumAuthorityVehicle string   `json:"num_authority_vehicle"`
	IsImportant       bool           `json:"is_important"`
	ManualMissionDate	string   `json:"manual_mission_date"`
	StationName       string   `json:"station_name"`
	AntaresId		  string   `json:"antares_id"`
	AntaresDescription string   `json:"antares_description"`
}

func ModelToMissionServiceSummaryJson(s models.MissionServiceSummary) *MissionServiceSummaryJson {
	service := MissionServiceSummaryJson{}

	service.Id = utils.ParseInt64String(s.Id)
	
	if s.Code.Valid {
		service.Code = s.Code.String
	}

	service.CreatedAt = s.ManualMissionDate.Time.String()

	service.Unharmed = utils.ParseInt64String(s.Unharmed.Int64)

	service.Injured = utils.ParseInt64String(s.Injured.Int64)

	service.Deceased = utils.ParseInt64String(s.Deceased.Int64)

	service.Transported = utils.ParseInt64String(s.Transported.Int64)

	service.AntaresDescription = s.AntaresDescription.String
	
	if s.NumServices.Valid {
		service.NumServices = utils.ParseInt64String(s.NumServices.Int64)
	}

	if s.NumAuthorityServices.Valid {
		service.NumAuthorityServices = utils.ParseInt64String(s.NumAuthorityServices.Int64)
	}

	if s.NumAuthorityPerson.Valid {
		service.NumAuthorityPerson = utils.ParseInt64String(s.NumAuthorityPerson.Int64)
	}

	if s.NumAuthorityVehicle.Valid {
		service.NumAuthorityVehicle = utils.ParseInt64String(s.NumAuthorityVehicle.Int64)
	}

	if s.NumServices.Valid {
		service.NumServices = utils.ParseInt64String(s.NumServices.Int64)
	}

	if s.AntaresId.Valid {
		service.AntaresId = utils.ParseInt64String(s.AntaresId.Int64)
	}

	if s.StationName.Valid {
		service.StationName = s.StationName.String
	}

	if s.NumVehicles.Valid {
		service.NumVehicles = utils.ParseInt64String(s.NumVehicles.Int64)
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


	service.Id = utils.ParseInt64(s.Id)

	service.AntaresId.Int64 = utils.ParseInt64(s.AntaresId)
	service.AntaresId.Valid = true

	service.StationName.String = s.StationName
	service.StationName.Valid = true

	serviceDate, err := time.Parse("02-01-2006 15:04:05", s.ManualMissionDate)
	if err == nil {
		service.ManualMissionDate.Time = serviceDate
		service.ManualMissionDate.Valid = true
	} else {
		logger.Warn().Err(err).Msg("Problema parseando service date")
	}

	manualServiceDate, err := time.Parse("02-01-2006 15:04:05", s.ManualMissionDate)
	if err == nil {
		service.ManualMissionDate.Time = manualServiceDate
		service.ManualMissionDate.Valid = true
	} else {
		logger.Warn().Err(err).Msg("Problema parseando manual service date")
	}

	service.NumVehicles.Int64 = utils.ParseInt64(s.NumVehicles)
	service.NumVehicles.Valid = true

	service.NumFirefighters.Int64 = utils.ParseInt64(s.NumFirefighters)
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
