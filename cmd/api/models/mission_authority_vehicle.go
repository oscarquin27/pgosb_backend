package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionAuthorityVehicleJson struct {
	Id          string `json:"id"`
	MissionId   string `json:"mission_id"`
	AuthorityId string `json:"authority_id"`
	Type        string `json:"type"`
	Make        string `json:"make"`
	Model       string `json:"model"`

	Plate       string `json:"plate"`
	Year        string `json:"year"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

func (m *MissionAuthorityVehicleJson) ToModel() models.MissionAuthorityVehicle {

	var missionAuthorityVehicle models.MissionAuthorityVehicle

	missionAuthorityVehicle.Id = utils.ParseInt64(m.Id)
	missionAuthorityVehicle.MissionId = utils.ParseInt64(m.MissionId)
	missionAuthorityVehicle.AuthorityId = utils.ParseInt64(m.AuthorityId)

	missionAuthorityVehicle.Type = sql.NullString{
		String: m.Type,
		Valid:  m.Type != "",
	}
	missionAuthorityVehicle.Make = sql.NullString{
		String: m.Make,
		Valid:  m.Make != "",
	}
	missionAuthorityVehicle.Model = sql.NullString{
		String: m.Model,
		Valid:  m.Model != "",
	}
	missionAuthorityVehicle.Plate = sql.NullString{
		String: m.Plate,
		Valid:  m.Plate != "",
	}
	missionAuthorityVehicle.Year = sql.NullString{
		String: m.Year,
		Valid:  m.Year != "",
	}
	missionAuthorityVehicle.Color = sql.NullString{
		String: m.Color,
		Valid:  m.Color != "",
	}
	missionAuthorityVehicle.Description = sql.NullString{
		String: m.Description,
		Valid:  m.Description != "",
	}

	return missionAuthorityVehicle

}

func ModelToMissionAuthorityVehicleJson(missionAuthorityVehicle *models.MissionAuthorityVehicle) *MissionAuthorityVehicleJson {

	m := &MissionAuthorityVehicleJson{}
	m.Id = utils.ParseInt64String(missionAuthorityVehicle.Id)
	m.MissionId = utils.ParseInt64String(missionAuthorityVehicle.MissionId)
	m.AuthorityId = utils.ParseInt64String(missionAuthorityVehicle.AuthorityId)

	if missionAuthorityVehicle.Type.Valid {
		m.Type = missionAuthorityVehicle.Type.String
	}
	if missionAuthorityVehicle.Make.Valid {
		m.Make = missionAuthorityVehicle.Make.String
	}
	if missionAuthorityVehicle.Model.Valid {
		m.Model = missionAuthorityVehicle.Model.String
	}
	if missionAuthorityVehicle.Plate.Valid {
		m.Plate = missionAuthorityVehicle.Plate.String
	}
	if missionAuthorityVehicle.Year.Valid {
		m.Year = missionAuthorityVehicle.Year.String
	}
	if missionAuthorityVehicle.Color.Valid {
		m.Color = missionAuthorityVehicle.Color.String
	}
	if missionAuthorityVehicle.Description.Valid {
		m.Description = missionAuthorityVehicle.Description.String
	}

	return m
}
