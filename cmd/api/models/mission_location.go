package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionLocationJson struct {
	Id string `json:"id"`

	MissionId string `json:"mission_id"`

	Alias string `json:"alias"`

	StateId string `json:"state_id"`
	State   string `json:"state"`

	MunicipalityId string `json:"municipality_id"`
	Municipality   string `json:"municipality"`

	ParishId string `json:"parish_id"`
	Parish   string `json:"parish"`

	SectorId string `json:"sector_id"`
	Sector   string `json:"sector"`

	UrbId string `json:"urb_id"`
	Urb   string `json:"urb"`

	// Street string `json:"street"`

	Address string `json:"address"`
}

func ModelToMissionLocationJson(s *models.MissionLocation) *MissionLocationJson {
	station := MissionLocationJson{
		Id:        utils.ParseInt64String(s.Id),
		MissionId: utils.ParseInt64String(s.MissionId),
		StateId:   utils.ParseInt64StringPointer(s.StateId),

		MunicipalityId: utils.ParseInt64StringPointer(s.MunicipalityId),
		ParishId:       utils.ParseInt64StringPointer(s.ParishId),
		SectorId:       utils.ParseInt64StringPointer(s.SectorId),
		UrbId:          utils.ParseInt64StringPointer(s.UrbId),

		Alias: utils.GetStringFromPointer(s.Alias),

		State:        utils.GetStringFromPointer(s.State),
		Municipality: utils.GetStringFromPointer(s.Municipality),
		Parish:       utils.GetStringFromPointer(s.Parish),
		Sector:       utils.GetStringFromPointer(s.Sector),
		Urb:          utils.GetStringFromPointer(s.Urb),
		// Street:       utils.GetStringFromPointer(s.Street),
		Address: utils.GetStringFromPointer(s.Address),
	}
	return &station
}

func (s *MissionLocationJson) ToModel() models.MissionLocation {

	state_id := utils.ParseInt64(s.StateId)
	municipality_id := utils.ParseInt64(s.MunicipalityId)
	parish_id := utils.ParseInt64(s.ParishId)
	sector_id := utils.ParseInt64(s.SectorId)
	urb_id := utils.ParseInt64(s.UrbId)

	station := models.MissionLocation{

		Id:        utils.ParseInt64(s.Id),
		MissionId: utils.ParseInt64(s.MissionId),

		StateId:        &state_id,
		MunicipalityId: &municipality_id,
		ParishId:       &parish_id,
		SectorId:       &sector_id,
		UrbId:          &urb_id,
		Alias:          &s.Alias,
		State:          &s.State,
		Municipality:   &s.Municipality,
		Parish:         &s.Parish,
		Sector:         &s.Sector,
		Urb:            &s.Urb,
		// Street:         &s.Street,
		Address: &s.Address,
	}
	return station
}
