package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionAuthorityJson struct {
	Id        string `json:"id"`
	MissionId string `json:"mission_id"`
	Alias     string `json:"alias"`
}

type MissionAuthoritySummaryJson struct {
	Id        string `json:"id"`
	MissionId string `json:"mission_id"`
	Alias     string `json:"alias"`
	Vehicles  string `json:"vehicles"`
	People    string `json:"people"`
	Type      string `json:"type"`
	Services  string `json:"services"`
}

func (m *MissionAuthorityJson) ToModel() models.MissionAuthority {

	model := models.MissionAuthority{}

	model.MissionId = utils.ParseInt64(m.MissionId)
	model.Id = utils.ParseInt64(m.Id)

	model.Alias = sql.NullString{

		String: m.Alias,
		Valid:  true,
	}

	return model
}

func (m *MissionAuthoritySummaryJson) ToModel() models.MissionAuthoritySummary {

	model := models.MissionAuthoritySummary{}

	model.MissionId = utils.ParseInt64(m.MissionId)
	model.Id = utils.ParseInt64(m.Id)
	model.Alias = sql.NullString{
		String: m.Alias,
		Valid:  true,
	}
	model.Services = sql.NullInt64{
		Int64: utils.ParseInt64(m.Services),
		Valid: true,
	}
	model.Vehicles = sql.NullInt64{

		Int64: utils.ParseInt64(m.Vehicles),
		Valid: true,
	}

	model.People = sql.NullInt64{
		Int64: utils.ParseInt64(m.People),
		Valid: true,
	}

	model.Type = sql.NullString{
		String: m.Type,
		Valid:  true,
	}

	return model
}

func ModelToMissionAuthorityJson(model *models.MissionAuthority) *MissionAuthorityJson {

	jsonModel := MissionAuthorityJson{}

	jsonModel.Id = utils.ParseInt64String(model.Id)
	jsonModel.MissionId = utils.ParseInt64String(model.MissionId)

	if model.Alias.Valid {
		jsonModel.Alias = model.Alias.String
	}

	return &jsonModel
}

func ModelToMissionAuthoritySummaryJson(model *models.MissionAuthoritySummary) *MissionAuthoritySummaryJson {

	jsonModel := MissionAuthoritySummaryJson{}

	jsonModel.Id = utils.ParseInt64String(model.Id)
	jsonModel.MissionId = utils.ParseInt64String(model.MissionId)

	if model.Alias.Valid {

		jsonModel.Alias = model.Alias.String
	}

	if model.Services.Valid {
		jsonModel.Services = utils.ParseInt64String(model.Services.Int64)
	}

	if model.Vehicles.Valid {

		jsonModel.Vehicles = utils.ParseInt64String(model.Vehicles.Int64)
	}

	if model.People.Valid {
		jsonModel.People = utils.ParseInt64String(model.People.Int64)
	}

	if model.Type.Valid {
		jsonModel.Type = model.Type.String
	}

	return &jsonModel
}
