package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type UrbanizationJson struct {
	Id       string `json:"id"`
	SectorId string `json:"sector_id"`
	Name     string `json:"name"`
}

func ModelToUrbanizationJson(s *models.Urbanization) *UrbanizationJson {
	state := UrbanizationJson{}

	state.Id = utils.ParseInt64Sring(s.Id)
	state.SectorId = utils.ParseInt64Sring(s.SectorId)
	state.Name = utils.GetStringFromPointer(s.Name)

	return &state
}

func (s *UrbanizationJson) ToModel() models.Urbanization {

	var state models.Urbanization

	state.Id = utils.ParseInt64(s.Id)
	state.SectorId = utils.ParseInt64(s.SectorId)
	state.Name = &s.Name

	return state
}
