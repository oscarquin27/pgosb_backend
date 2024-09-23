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

	unit.Id = utils.ParseInt64String(s.Id)
	unit.Plate = utils.GetStringFromPointer(s.Plate)
	unit.Station = utils.GetStringFromPointer(s.Station)

	unit.Unit_type = utils.GetStringFromPointer(s.Unit_type)
	unit.Make = utils.GetStringFromPointer(s.Make)
	unit.Unit_condition = utils.GetStringFromPointer(s.Unit_condition)
	unit.Vehicle_serial = utils.GetStringFromPointer(s.Vehicle_serial)
	unit.Motor_serial = utils.GetStringFromPointer(s.Motor_serial)
	unit.Capacity = utils.GetStringFromPointer(s.Capacity)
	unit.Fuel_type = utils.GetStringFromPointer(s.Fuel_type)
	unit.Water_capacity = utils.GetStringFromPointer(s.Water_capacity)
	unit.Observations = utils.GetStringFromPointer(s.Observations)

	// Direct Assignment for String Slices
	unit.Details = s.Details

	// Conversion for Additional Fields
	unit.Hurt_capacity = utils.ParseInt64StringPointer(s.Hurt_capacity)
	unit.Doors = utils.ParseInt64StringPointer(s.Doors)
	unit.Performance = utils.GetStringFromPointer(s.Performance)
	unit.Load_capacity = utils.ParseInt64StringPointer(s.Load_capacity)

	unit.Model = utils.GetStringFromPointer(s.Model)
	unit.Alias = utils.GetStringFromPointer(s.Alias)
	unit.Color = utils.GetStringFromPointer(s.Color)
	unit.Year = utils.GetStringFromPointer(s.Year)
	unit.Purpose = utils.GetStringFromPointer(s.Purpose)
	unit.Init_kilometer = utils.ParseInt64StringPointer(s.Init_kilometer)

	return &unit

}

func (s *UnitJson) ToModel() models.Unit {
	var unit models.Unit

	// Basic String Conversions
	unit.Id = utils.ParseInt64(s.Id)
	unit.Plate = &s.Plate
	unit.Station = &s.Station
	unit.Unit_type = &s.Unit_type
	unit.Make = &s.Make
	unit.Unit_condition = &s.Unit_condition
	unit.Vehicle_serial = &s.Vehicle_serial
	unit.Motor_serial = &s.Motor_serial
	unit.Capacity = &s.Capacity
	unit.Fuel_type = &s.Fuel_type
	unit.Water_capacity = &s.Water_capacity
	unit.Observations = &s.Observations
	unit.Details = s.Details // Direct assignment for string slices

	// Conversion for Additional Fields

	hurt_capacity := utils.ParseInt64(s.Hurt_capacity)
	unit.Hurt_capacity = &hurt_capacity

	doors := utils.ParseInt64(s.Doors)
	unit.Doors = &doors

	unit.Performance = &s.Performance

	load_capacity := utils.ParseInt64(s.Load_capacity)
	unit.Load_capacity = &load_capacity

	unit.Model = &s.Model
	unit.Alias = &s.Alias
	unit.Color = &s.Color
	unit.Year = &s.Year
	unit.Purpose = &s.Purpose

	init_kilometer := utils.ParseInt64(s.Init_kilometer)
	unit.Init_kilometer = &init_kilometer

	return unit
}
