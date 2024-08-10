package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type StateJson struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Area_Code   string `json:"area_code"`
	Zip_Code    string `json:"zip_code"`
	Coordinates string `json:"Coordinates"`
}

func ModelToStateJson(s *models.State) *StateJson {
	state := StateJson{}

	state.Id = utils.ParseInt64Sring(s.Id)
	state.Name = s.Name
	state.Area_Code = s.Area_Code
	state.Zip_Code = s.Zip_Code
	state.Coordinates = s.Coordinates

	return &state
}

func (s *StateJson) ToModel() models.State {

	var state models.State

	state.Id = utils.ParseInt64(s.Id)

	state.Name = s.Name
	state.Area_Code = s.Area_Code
	state.Zip_Code = s.Zip_Code
	state.Coordinates = s.Coordinates

	return state
}
