package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type UnitJson struct {
	Id             string   `json:"id"`
	Plate          string   `json:"plate"`
	Station        string   `json:"station"`
	Unit_type      string   `json:"unit_type"`
	Make           string   `json:"make"`
	Unit_condition string   `json:"unit_condition"`
	Vehicle_serial string   `json:"vehicle_serial"`
	Motor_serial   string   `json:"motor_serial"`
	Capacity       string   `json:"capacity"`
	Details        []string `json:"details"`
	Fuel_type      string   `json:"fuel_type"`
	Water_capacity string   `json:"water_capacity"`
	Observations   string   `json:"observations"`
	Hurt_capacity  string   `json:"hurt_capacity"`
	Doors          string   `json:"doors"`
	Performance    string   `json:"performance"`
	Load_capacity  string   `json:"load_capacity"`
	Model          string   `json:"model"`
	Alias          string   `json:"alias"`
	Color          string   `json:"color"`
	Year           string   `json:"year"`
	Purpose        string   `json:"purpose"`
	Init_kilometer string   `json:"init_kilometer"`
}

func ModelToUnitJson(s *models.Unit) *UnitJson {
	unit := UnitJson{}

	// Basic String Conversions
	unit.Id = utils.ConvertFromInt4(s.Id)
	unit.Plate = utils.ConvertFromText(s.Plate)
	unit.Station = utils.ConvertFromText(s.Station)
	unit.Unit_type = utils.ConvertFromText(s.Unit_type)
	unit.Make = utils.ConvertFromText(s.Make)
	unit.Unit_condition = utils.ConvertFromText(s.Unit_condition)
	unit.Vehicle_serial = utils.ConvertFromText(s.Vehicle_serial)
	unit.Motor_serial = utils.ConvertFromText(s.Motor_serial)
	unit.Capacity = utils.ConvertFromText(s.Capacity)
	unit.Fuel_type = utils.ConvertFromText(s.Fuel_type)
	unit.Water_capacity = utils.ConvertFromText(s.Water_capacity)
	unit.Observations = utils.ConvertFromText(s.Observations)

	// Direct Assignment for String Slices
	unit.Details = s.Details

	// Conversion for Additional Fields
	unit.Hurt_capacity = utils.ConvertFromInt4(s.Hurt_capacity)
	unit.Doors = utils.ConvertFromInt4(s.Doors)
	unit.Performance = utils.ConvertFromText(s.Performance)
	unit.Load_capacity = utils.ConvertFromInt4(s.Load_capacity)
	unit.Model = utils.ConvertFromText(s.Model)
	unit.Alias = utils.ConvertFromText(s.Alias)
	unit.Color = utils.ConvertFromText(s.Color)
	unit.Year = utils.ConvertFromText(s.Year)
	unit.Purpose = utils.ConvertFromText(s.Purpose)
	unit.Init_kilometer = utils.ConvertFromInt4(s.Init_kilometer)

	return &unit
}

func (s *UnitJson) ToModel() models.Unit {
	var unit models.Unit

	// Basic String Conversions
	unit.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	unit.Plate = utils.ConvertToPgTypeText(s.Plate)
	unit.Station = utils.ConvertToPgTypeText(s.Station)
	unit.Unit_type = utils.ConvertToPgTypeText(s.Unit_type)
	unit.Make = utils.ConvertToPgTypeText(s.Make)
	unit.Unit_condition = utils.ConvertToPgTypeText(s.Unit_condition)
	unit.Vehicle_serial = utils.ConvertToPgTypeText(s.Vehicle_serial)
	unit.Motor_serial = utils.ConvertToPgTypeText(s.Motor_serial)
	unit.Capacity = utils.ConvertToPgTypeText(s.Capacity)
	unit.Fuel_type = utils.ConvertToPgTypeText(s.Fuel_type)
	unit.Water_capacity = utils.ConvertToPgTypeText(s.Water_capacity)
	unit.Observations = utils.ConvertToPgTypeText(s.Observations)
	unit.Details = s.Details // Direct assignment for string slices

	// Conversion for Additional Fields
	unit.Hurt_capacity = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Hurt_capacity))
	unit.Doors = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Doors))
	unit.Performance = utils.ConvertToPgTypeText(s.Performance)
	unit.Load_capacity = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Load_capacity))
	unit.Model = utils.ConvertToPgTypeText(s.Model)
	unit.Alias = utils.ConvertToPgTypeText(s.Alias)
	unit.Color = utils.ConvertToPgTypeText(s.Color)
	unit.Year = utils.ConvertToPgTypeText(s.Year)
	unit.Purpose = utils.ConvertToPgTypeText(s.Purpose)
	unit.Init_kilometer = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Init_kilometer))

	return unit
}
