package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MunicipalityJson struct {
	Id          string `json:"id"`
	StateId     string `json:"state_id"`
	Name        string `json:"name"`
	Capital     string `json:"capital"`
	Coordinates string `json:"coordinates"`
}

func ModelToMunicipalityJson(s *models.Municipality) *MunicipalityJson {
	state := MunicipalityJson{}

	state.Id = utils.ParseInt64String(s.Id)
	state.StateId = utils.ParseInt64String(s.StateId)

	state.Name = utils.GetStringFromPointer(s.Name)
	state.Capital = utils.GetStringFromPointer(s.Capital)
	state.Coordinates = utils.GetStringFromPointer(s.Coordinates)

	return &state
}

func (s *MunicipalityJson) ToModel() models.Municipality {

	var state models.Municipality

	state.Id = utils.ParseInt64(s.Id)
	state.StateId = utils.ParseInt64(s.StateId)
	state.Name = &s.Name
	state.Capital = &s.Capital
	state.Coordinates = &s.Coordinates

	return state
}
