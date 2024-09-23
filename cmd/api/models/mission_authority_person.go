package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionAuthorityPersonJson struct {
	Id                   string `json:"id"`
	MissionId            string `json:"mission_id"`
	AuthorityId          string `json:"authority_id"`
	Name                 string `json:"name"`
	LastName             string `json:"last_name"`
	LegalId              string `json:"legal_id"`
	IdentificationNumber string `json:"identification_number"`
	Phone                string `json:"phone"`
	Gender               string `json:"gender"`
	Observations         string `json:"observations"`
}

func (m *MissionAuthorityPersonJson) ToModel() models.MissionAuthorityPerson {
	model := models.MissionAuthorityPerson{}

	model.Id = utils.ParseInt64(m.Id)
	model.MissionId = utils.ParseInt64(m.MissionId)
	model.AuthorityId = utils.ParseInt64(m.AuthorityId)
	model.Name = sql.NullString{
		String: m.Name,
		Valid:  true,
	}
	model.LastName = sql.NullString{
		String: m.LastName,
		Valid:  true,
	}
	model.LegalId = sql.NullString{
		String: m.LegalId,
		Valid:  true,
	}
	model.IdentificationNumber = sql.NullString{
		String: m.IdentificationNumber,
		Valid:  true,
	}
	model.Phone = sql.NullString{
		String: m.Phone,
		Valid:  true,
	}
	model.Gender = sql.NullString{
		String: m.Gender,
		Valid:  true,
	}
	model.Observations = sql.NullString{
		String: m.Observations,
		Valid:  true,
	}

	return model
}

func ModelToMissionAuthorityPersonJson(model *models.MissionAuthorityPerson) *MissionAuthorityPersonJson {
	jsonModel := MissionAuthorityPersonJson{}

	jsonModel.Id = utils.ParseInt64String(model.Id)
	jsonModel.MissionId = utils.ParseInt64String(model.MissionId)
	jsonModel.AuthorityId = utils.ParseInt64String(model.AuthorityId)

	if model.Name.Valid {
		jsonModel.Name = model.Name.String
	}
	if model.LastName.Valid {
		jsonModel.LastName = model.LastName.String
	}
	if model.LegalId.Valid {
		jsonModel.LegalId = model.LegalId.String
	}
	if model.IdentificationNumber.Valid {
		jsonModel.IdentificationNumber = model.IdentificationNumber.String
	}
	if model.Phone.Valid {
		jsonModel.Phone = model.Phone.String
	}
	if model.Gender.Valid {
		jsonModel.Gender = model.Gender.String
	}
	if model.Observations.Valid {
		jsonModel.Observations = model.Observations.String
	}

	return &jsonModel
}
