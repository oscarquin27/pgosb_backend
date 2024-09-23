package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type AuthorityJson struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Government   string `json:"government"`
}

func (a *AuthorityJson) ToModel() models.Authority {

	var authority models.Authority

	authority.Id = utils.ParseInt64(a.Id)

	authority.Name = sql.NullString{
		String: a.Name,
		Valid:  a.Name != "",
	}

	authority.Abbreviation = sql.NullString{
		String: a.Abbreviation,
		Valid:  a.Abbreviation != "",
	}

	authority.Government = sql.NullString{
		String: a.Government,
		Valid:  a.Government != "",
	}

	return authority
}

func ModelToAuthorityJson(model *models.Authority) *AuthorityJson {

	jsonModel := AuthorityJson{}

	jsonModel.Id = utils.ParseInt64String(model.Id)

	if model.Name.Valid {
		jsonModel.Name = model.Name.String
	}
	if model.Abbreviation.Valid {
		jsonModel.Abbreviation = model.Abbreviation.String
	}
	if model.Government.Valid {
		jsonModel.Government = model.Government.String
	}

	return &jsonModel

}
