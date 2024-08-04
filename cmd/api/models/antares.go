package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
	"strconv"
)

type AntaresJson struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func ModelToAntaresJson(a models.Antares) *AntaresJson {
	antaresJson := &AntaresJson{}

	antaresJson.Id = strconv.Itoa(a.Id) // Assuming utils.ConvertFromInt4 converts int32 to string
	antaresJson.Type = a.Type
	antaresJson.Description = a.Description

	return antaresJson
}

func (a *AntaresJson) ToModel() models.Antares {
	antares := models.Antares{}

	antares.Id = utils.ParseInt(a.Id) // Assuming utils.ConvertToPgTypeInt4 converts string to int32
	antares.Type = a.Type
	antares.Description = a.Description

	return antares
}
