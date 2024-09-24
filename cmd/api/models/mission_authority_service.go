package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionAuthorityServiceJson struct {
	Id          string `json:"id"`
	MissionId   string `json:"mission_id"`
	ServiceId   string `json:"service_id"`
	AuthorityId string `json:"authority_id"`
}

func (s *MissionAuthorityServiceJson) ToModel() models.MissionAuthorityService {
	var m models.MissionAuthorityService

	m.Id = utils.ParseInt64(s.Id)
	m.MissionId = utils.ParseInt64(s.MissionId)
	m.ServiceId = utils.ParseInt64(s.ServiceId)
	m.AuthorityId = utils.ParseInt64(s.AuthorityId)

	return m
}

func ModelToMissionAuthorityServiceJson(model *models.MissionAuthorityService) *MissionAuthorityServiceJson {
	var s MissionAuthorityServiceJson

	s.Id = utils.ParseInt64String(model.Id)
	s.MissionId = utils.ParseInt64String(model.MissionId)
	s.ServiceId = utils.ParseInt64String(model.ServiceId)
	s.AuthorityId = utils.ParseInt64String(model.AuthorityId)

	return &s
}

type MissionAuthorityServiceSummaryJson struct {
	Id          string `json:"id"`
	AuthorityId string `json:"authority_id"`
	MissionId   string `json:"mission_id"`

	//CreatedAt string   `json:"created_at"`
	Alias    string `json:"alias"`
	Vehicles string `json:"vehicles"`
	People   string `json:"people"`
	Type     string `json:"type"`
	Services string `json:"services"`
}

func (s *MissionAuthorityServiceSummaryJson) ToModel() models.MissionAuthorityServiceSummary {
	var m models.MissionAuthorityServiceSummary

	m.Id = utils.ParseInt64(s.Id)
	m.AuthorityId = utils.ParseInt64(s.AuthorityId)
	m.MissionId = utils.ParseInt64(s.MissionId)

	m.Alias = sql.NullString{
		String: s.Alias,
		Valid:  s.Alias != "",
	}
	m.Vehicles = sql.NullInt64{
		Int64: utils.ParseInt64(s.Vehicles),
		Valid: s.Vehicles != "",
	}
	m.People = sql.NullInt64{
		Int64: utils.ParseInt64(s.People),
		Valid: s.People != "",
	}
	m.Type = sql.NullString{
		String: s.Type,
		Valid:  s.Type != "",
	}

	m.Services = sql.NullInt64{
		Int64: utils.ParseInt64(s.Services),
		Valid: s.Services != "",
	}

	return m
}

func ModelToMissionAuthorityServiceSummaryJson(model *models.MissionAuthorityServiceSummary) *MissionAuthorityServiceSummaryJson {
	var s MissionAuthorityServiceSummaryJson

	s.Id = utils.ParseInt64String(model.Id)
	s.AuthorityId = utils.ParseInt64String(model.AuthorityId)
	s.MissionId = utils.ParseInt64String(model.MissionId)

	if model.Alias.Valid {
		s.Alias = model.Alias.String
	}
	if model.Vehicles.Valid {
		s.Vehicles = utils.ParseInt64String(model.Vehicles.Int64)
	}
	if model.People.Valid {
		s.People = utils.ParseInt64String(model.People.Int64)
	}
	if model.Type.Valid {
		s.Type = model.Type.String
	}
	if model.Services.Valid {
		s.Services = utils.ParseInt64String(model.Services.Int64)
	}

	return &s
}
