package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type ParishJson struct {
	Id             string `json:"id"`
	StateId        string `json:"state_id"`
	MunicipalityId string `json:"municipality_id"`
	Name           string `json:"name"`
	Capital        string `json:"capital"`
	Coordinates    string `json:"coordinates"`
}

func ModelToParishJson(s *models.Parish) *ParishJson {
	state := ParishJson{}

	state.Id = utils.ParseInt64String(s.Id)
	state.StateId = utils.ParseInt64String(s.StateId)
	state.MunicipalityId = utils.ParseInt64String(s.MunicipalityId)

	state.Name = utils.GetStringFromPointer(s.Name)
	state.Capital = utils.GetStringFromPointer(s.Capital)
	state.Coordinates = utils.GetStringFromPointer(s.Coordinates)

	return &state
}

func (s *ParishJson) ToModel() models.Parish {

	var state models.Parish

	state.Id = utils.ParseInt64(s.Id)
	state.StateId = utils.ParseInt64(s.StateId)
	state.MunicipalityId = utils.ParseInt64(s.MunicipalityId)

	state.Name = &s.Name
	state.Capital = &s.Capital
	state.Coordinates = &s.Coordinates

	return state
}
