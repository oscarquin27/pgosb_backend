package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type OperativeRegionsJson struct {
	Id           string `json:"id"`
	Description  string `json:"description"`
	Abbreviation string `json:"abbreviation"`
	Phone        string `json:"phone"`
	Coverage	 string `json:"coverage"`
}

func ModelToOperativeRegionsJson(s *models.OperativeRegions) *OperativeRegionsJson {
	or := OperativeRegionsJson{}

	or.Id = utils.ParseInt64Sring(s.Id)
	or.Description = utils.GetStringFromPointer(s.Description)
	or.Abbreviation = utils.GetStringFromPointer(s.Abbreviation)
	or.Phone = utils.GetStringFromPointer(s.Phone)
	or.Coverage = utils.GetStringFromPointer(s.Coverage)


	return &or
}

func (s *OperativeRegionsJson) ToModel() models.OperativeRegions {

	var or models.OperativeRegions

	or.Id = utils.ParseInt64(s.Id)
	or.Description = &s.Description
	or.Abbreviation = &s.Abbreviation
	or.Phone = &s.Phone
	or.Coverage = &s.Coverage

	return or
}
