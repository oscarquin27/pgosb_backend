package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type UnitJson struct {
	Id             string `json:"id"`
	Plate          string `json:"plate"`
	Zone           string `json:"zone"`
	Station        string `json:"station"`
	Unit_type      string `json:"unit_type"`
	Make           string `json:"make"`
	Drivers        string `json:"drivers"`
	Unit_condition string `json:"unit_condition"`
	Vehicle_serial string `json:"vehicle_serial"`
	Motor_serial   string `json:"motor_serial"`
	Capacity       string `json:"capacity"`
	Details        string `json:"details"`
	Fuel_type      string `json:"fuel_type"`
	Water_capacity string `json:"water_capacity"`
	Observations   string `json:"observations"`
}

func (s *UnitJson) ToModel() models.Unit {
	var unit models.Unit

	unit.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	unit.Plate = utils.ConvertToPgTypeText(s.Plate)
	unit.Zone = utils.ConvertToPgTypeText(s.Zone)
	unit.Station = utils.ConvertToPgTypeText(s.Station)
	unit.Unit_type = utils.ConvertToPgTypeText(s.Unit_type)
	unit.Make = utils.ConvertToPgTypeText(s.Make)
	unit.Drivers = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Drivers))
	unit.Unit_condition = utils.ConvertToPgTypeText(s.Unit_condition)
	unit.Vehicle_serial = utils.ConvertToPgTypeText(s.Vehicle_serial)
	unit.Motor_serial = utils.ConvertToPgTypeText(s.Motor_serial)
	unit.Capacity = utils.ConvertToPgTypeText(s.Capacity)
	unit.Fuel_type = utils.ConvertToPgTypeText(s.Fuel_type)
	unit.Water_capacity = utils.ConvertToPgTypeText(s.Water_capacity)
	unit.Observations = utils.ConvertToPgTypeText(s.Observations)

	return unit
}
