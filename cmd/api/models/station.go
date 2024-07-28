package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type StationJson struct {
	Id              string    `json:"id" db:"id"`
	Municipality_id string    `json:"municipality_id"`
	Name            string    `json:"name"`
	Coordinates     string    `json:"coordinates"`
	Description     string    `json:"description"`
	Code            string    `json:"code"`
	Abbreviation    string    `json:"abbreviation"`
	Phones          *[]string `json:"phones"`
	Regions         *[]string `json:"regions"`
	State_id        string    `json:"state_id"`
	Parish_id       string    `json:"parish_id"`
	Sector          string    `json:"sector"`
	Community       string    `json:"community"`
	Street          string    `json:"street"`
	Institution     string    `json:"institution"`
	State           string    `json:"state"`
	Municipality    string    `json:"municipality"`
	Parish          string    `json:"parish"`
	Address         string    `json:"address"`
}

func ModelToStationJson(s models.Station) *StationJson {
	station := StationJson{
		Id:              utils.ConvertFromInt4(s.Id),
		Municipality_id: utils.ConvertFromInt4(s.Municipality_id),
		Name:            utils.ConvertFromText(s.Name),
		Coordinates:     utils.ConvertFromText(s.Coordinates),
		Description:     utils.ConvertFromText(s.Description),
		Code:            utils.ConvertFromText(s.Code),
		Abbreviation:    utils.ConvertFromText(s.Abbreviation),
		Phones:          s.Phones,
		Regions:         s.Regions, // Make sure this is included if you're using it
		State_id:        utils.ConvertFromInt4(s.State_id),
		Parish_id:       utils.ConvertFromInt4(s.Parish_id),
		Sector:          utils.ConvertFromText(s.Sector),
		Community:       utils.ConvertFromText(s.Community),
		Street:          utils.ConvertFromText(s.Street),
		Institution:     utils.ConvertFromText(s.Institution),
		State:           utils.ConvertFromText(s.State),
		Municipality:    utils.ConvertFromText(s.Municipality),
		Parish:          utils.ConvertFromText(s.Parish),
		Address:         utils.ConvertFromText(s.Address),
	}
	return &station
}

func (s *StationJson) ToModel() models.Station {
	station := models.Station{
		Id:              utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id)),
		Municipality_id: utils.ConvertToPgTypeInt4(utils.ParseInt(s.Municipality_id)),
		Name:            utils.ConvertToPgTypeText(s.Name),
		Coordinates:     utils.ConvertToPgTypeText(s.Coordinates),
		Description:     utils.ConvertToPgTypeText(s.Description),
		Code:            utils.ConvertToPgTypeText(s.Code),
		Abbreviation:    utils.ConvertToPgTypeText(s.Abbreviation),
		Phones:          s.Phones,
		Regions:         s.Regions, // Make sure this is included if you're using it
		State_id:        utils.ConvertToPgTypeInt4(utils.ParseInt(s.State_id)),
		Parish_id:       utils.ConvertToPgTypeInt4(utils.ParseInt(s.Parish_id)),
		Sector:          utils.ConvertToPgTypeText(s.Sector),
		Community:       utils.ConvertToPgTypeText(s.Community),
		Street:          utils.ConvertToPgTypeText(s.Street),
		Institution:     utils.ConvertToPgTypeText(s.Institution),
		State:           utils.ConvertToPgTypeText(s.State),
		Municipality:    utils.ConvertToPgTypeText(s.Municipality),
		Parish:          utils.ConvertToPgTypeText(s.Parish),
		Address:         utils.ConvertToPgTypeText(s.Address),
	}
	return station
}
