package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionVehicleJson struct {
	Id               string `json:"id"`
	MissionId        string `json:"mission_id"`
	UnitId           string `json:"unit_id"`
	VehicleCondition string `json:"vehicle_condition"`
	Make             string `json:"make"`
	Model            string `json:"model"`
	Year             string `json:"year"`
	MotorSerial      string `json:"motor_serial"`
	Plate            string `json:"plate"`
	Color            string `json:"color"`
	VehicleType      string `json:"vehicle_type"`
	VehicleVerified  bool   `json:"vehicle_verified"`
}

func ModelToMissionVehicleJson(s models.MissionVehicle) *MissionVehicleJson {
	vehicle := MissionVehicleJson{}

	vehicle.Id = utils.ConvertFromInt4(s.Id)
	vehicle.MissionId = utils.ConvertFromInt4(s.MissionId)
	vehicle.VehicleCondition = utils.ConvertFromText(s.VehicleCondition)
	vehicle.Make = utils.ConvertFromText(s.Make)
	vehicle.Model = utils.ConvertFromText(s.Model)
	vehicle.Year = utils.ConvertFromText(s.Year)
	vehicle.Plate = utils.ConvertFromText(s.Plate)
	vehicle.Color = utils.ConvertFromText(s.Color)
	vehicle.VehicleType = utils.ConvertFromText(s.VehicleType)
	vehicle.MotorSerial = utils.ConvertFromText(s.MotorSerial)
	vehicle.VehicleVerified = utils.ConvertFromBool(s.VehicleVerified)

	return &vehicle
}

func (s *MissionVehicleJson) ToModel() models.MissionVehicle {
	vehicle := models.MissionVehicle{}

	vehicle.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	vehicle.MissionId = utils.ConvertToPgTypeInt4(utils.ParseInt(s.MissionId))
	vehicle.VehicleCondition = utils.ConvertToPgTypeText(s.VehicleCondition)
	vehicle.Make = utils.ConvertToPgTypeText(s.Make)
	vehicle.Model = utils.ConvertToPgTypeText(s.Model)
	vehicle.Year = utils.ConvertToPgTypeText(s.Year)
	vehicle.Plate = utils.ConvertToPgTypeText(s.Plate)
	vehicle.Color = utils.ConvertToPgTypeText(s.Color)
	vehicle.VehicleType = utils.ConvertToPgTypeText(s.VehicleType)
	vehicle.MotorSerial = utils.ConvertToPgTypeText(s.MotorSerial)
	vehicle.VehicleVerified = utils.ConvertToPgTypeBool(s.VehicleVerified)

	return vehicle
}
