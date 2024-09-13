package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionAuthorityJson struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	Name            string `json:"name"`
	Lastname        string `json:"lastname"`
	Identification  string `json:"identification"`
	Rank            string `json:"rank"`
	Phone           string `json:"phone"`
	Vehicles        string `json:"vehicles"`
	DetailOfVehicle string `json:"detail_of_vehicle"`
}

func (m *MissionAuthorityJson) ToModel() models.MissionAuthority {
	model := models.MissionAuthority{}

	model.Id = sql.NullInt64{
		Int64: utils.ParseInt64(m.Id),
		Valid: true,
	}

	model.Type = sql.NullString{
		String: m.Type,
		Valid:  true,
	}

	model.Name = sql.NullString{
		String: m.Name,
		Valid:  true,
	}

	model.Lastname = sql.NullString{
		String: m.Lastname,
		Valid:  true,
	}

	model.Identification = sql.NullString{
		String: m.Identification,
		Valid:  true,
	}

	model.Phone = sql.NullString{
		String: m.Phone,
		Valid:  true,
	}

	model.NumberVehicle = sql.NullString{
		String: m.Vehicles,
		Valid:  true,
	}

	model.DetailOfVehicle = sql.NullString{
		String: m.DetailOfVehicle,
		Valid:  true,
	}

	model.Rank = sql.NullString{
		String: m.Rank,
		Valid:  true,
	}

	return model
}

func ModelToMissionAuthorityJson(model *models.MissionAuthority) *MissionAuthorityJson {

	jsonModel := MissionAuthorityJson{}

	if model.Id.Valid {
		jsonModel.Id = utils.ParseInt64Sring(model.Id.Int64)
	}

	if model.Type.Valid {
		jsonModel.Type = model.Type.String
	}

	if model.Name.Valid {
		jsonModel.Name = model.Name.String
	}

	if model.Lastname.Valid {
		jsonModel.Lastname = model.Lastname.String
	}

	if model.Identification.Valid {
		jsonModel.Identification = model.Identification.String
	}

	if model.Phone.Valid {
		jsonModel.Phone = model.Phone.String
	}

	if model.NumberVehicle.Valid {
		jsonModel.Vehicles = model.NumberVehicle.String
	}

	if model.DetailOfVehicle.Valid {
		jsonModel.DetailOfVehicle = model.DetailOfVehicle.String
	}

	if model.Rank.Valid {
		jsonModel.Rank = model.Rank.String
	}

	return &jsonModel

}
