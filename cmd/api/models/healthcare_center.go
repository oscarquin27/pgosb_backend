package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type HealthcareCenterJson struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Abbreviation string   `json:"abbreviation"`
	Phones       []string `json:"phones"`
	RegionId     string   `json:"region_id"`

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

	Street string `json:"street"`

	Address string `json:"address"`
}

func ModelToHealthcareCenterJson(s *models.HealthcareCenter) *HealthcareCenterJson {
	station := HealthcareCenterJson{
		Id:             utils.ParseInt64Sring(s.Id),
		RegionId:       utils.ParseInt64SringPointer(s.RegionId),
		StateId:        utils.ParseInt64SringPointer(s.StateId),
		MunicipalityId: utils.ParseInt64SringPointer(s.MunicipalityId),
		ParishId:       utils.ParseInt64SringPointer(s.ParishId),
		SectorId:       utils.ParseInt64SringPointer(s.SectorId),
		UrbId:          utils.ParseInt64SringPointer(s.UrbId),
		Name:           utils.GetStringFromPointer(s.Name),
		Description:    utils.GetStringFromPointer(s.Description),
		Abbreviation:   utils.GetStringFromPointer(s.Abbreviation),
		Phones:         s.Phones,
		State:          utils.GetStringFromPointer(s.State),
		Municipality:   utils.GetStringFromPointer(s.Municipality),
		Parish:         utils.GetStringFromPointer(s.Parish),
		Sector:         utils.GetStringFromPointer(s.Sector),
		Urb:            utils.GetStringFromPointer(s.Urb),
		Street:         utils.GetStringFromPointer(s.Street),
		Address:        utils.GetStringFromPointer(s.Address),
	}
	return &station
}

func (s *HealthcareCenterJson) ToModel() models.HealthcareCenter {

	regionId := utils.ParseInt64(s.RegionId)
	state_id := utils.ParseInt64(s.StateId)
	municipality_id := utils.ParseInt64(s.MunicipalityId)
	parish_id := utils.ParseInt64(s.ParishId)
	sector_id := utils.ParseInt64(s.SectorId)
	urb_id := utils.ParseInt64(s.UrbId)

	station := models.HealthcareCenter{

		Id: utils.ParseInt64(s.Id),

		RegionId:       &regionId,
		StateId:        &state_id,
		MunicipalityId: &municipality_id,
		ParishId:       &parish_id,
		SectorId:       &sector_id,
		UrbId:          &urb_id,
		Name:           &s.Name,
		Description:    &s.Description,
		Abbreviation:   &s.Abbreviation,
		Phones:         s.Phones,
		State:          &s.State,
		Municipality:   &s.Municipality,
		Parish:         &s.Parish,
		Sector:         &s.Sector,
		Urb:            &s.Urb,
		Street:         &s.Street,
		Address:        &s.Address,
	}
	return station
}
