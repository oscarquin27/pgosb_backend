package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type StationJson struct {
	Id              string          `json:"id"`
	Municipality_id string          `json:"municipality_id"`
	Name            string          `json:"name"`
	Coordinates     string          `json:"coordinates"`
	Description     string          `json:"description"`
	Code            string          `json:"code"`
	Abbreviation    string          `json:"abbreviation"`
	Phones          []models.Phones `json:"phones"`
	State_id        string          `json:"state_id"`
	Parish_id       string          `json:"parish_id"`
	Sector          string          `json:"sector"`
	Community       string          `json:"community"`
	Street          string          `json:"street"`
	Address         string          `json:"address"`
}

func ModelToStationJson(s models.Station) *StationJson {
	station := StationJson{}

	station.Municipality_id = utils.ConvertFromInt4(s.Municipality_id)
	station.Name = utils.ConvertFromText(s.Name)
	station.Coordinates = utils.ConvertFromText(s.Coordinates)
	station.Description = utils.ConvertFromText(s.Description)
	station.Code = utils.ConvertFromText(s.Code)
	station.Abbreviation = utils.ConvertFromText(s.Abbreviation)
	station.Phones = s.Phones
	station.State_id = utils.ConvertFromInt4(s.State_id)
	station.Parish_id = utils.ConvertFromInt4(s.Parish_id)
	station.Sector = utils.ConvertFromText(s.Sector)
	station.Community = utils.ConvertFromText(s.Community)
	station.Street = utils.ConvertFromText(s.Street)
	station.Address = utils.ConvertFromText(s.Address)

	return &station
}

func (s *StationJson) ToModel() models.Station {
	station := models.Station{}

	station.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	station.Municipality_id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Municipality_id))
	station.Name = utils.ConvertToPgTypeText(s.Name)
	station.Coordinates = utils.ConvertToPgTypeText(s.Coordinates)
	station.Description = utils.ConvertToPgTypeText(s.Description)
	station.Code = utils.ConvertToPgTypeText(s.Code)
	station.Abbreviation = utils.ConvertToPgTypeText(s.Abbreviation)
	station.Phones = s.Phones
	station.State_id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.State_id))
	station.Parish_id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Parish_id))
	station.Sector = utils.ConvertToPgTypeText(s.Sector)
	station.Community = utils.ConvertToPgTypeText(s.Community)
	station.Street = utils.ConvertToPgTypeText(s.Street)
	station.Address = utils.ConvertToPgTypeText(s.Address)

	return station
}
