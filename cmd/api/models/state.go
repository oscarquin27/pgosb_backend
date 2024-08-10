package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type StateJson struct {
	Id          string `json:"id"`
	Name        string `db:"name"`
	Capital     string `db:"capital"`
	Coordinates string `db:"coordinates"`
}

func ModelToStateJson(s *models.State) *StateJson {
	state := StateJson{}

	state.Id = utils.ParseInt64Sring(s.Id)
	state.Name = utils.GetStringFromPointer(s.Name)
	state.Capital = utils.GetStringFromPointer(s.Capital)
	state.Coordinates = utils.GetStringFromPointer(s.Coordinates)

	return &state
}

func (s *StateJson) ToModel() models.State {

	var state models.State

	state.Id = utils.ParseInt64(s.Id)

	state.Name = &s.Name
	state.Capital = &s.Capital
	state.Coordinates = &s.Coordinates

	return state
}
