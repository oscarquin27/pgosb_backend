package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type SectorJson struct {
	Id       string `json:"id"`
	ParishId string `json:"parish_id"`
	Name     string `json:"name"`
}

func ModelToSectorJson(s *models.Sector) *SectorJson {
	state := SectorJson{}

	state.Id = utils.ParseInt64String(s.Id)
	state.ParishId = utils.ParseInt64String(s.ParishId)
	state.Name = utils.GetStringFromPointer(s.Name)

	return &state
}

func (s *SectorJson) ToModel() models.Sector {

	var state models.Sector

	state.Id = utils.ParseInt64(s.Id)
	state.ParishId = utils.ParseInt64(s.ParishId)
	state.Name = &s.Name

	return state
}
