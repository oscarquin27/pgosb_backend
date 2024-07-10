package unit_entity

type Unit struct {
	Id             int    `json:"id"`
	Plate          string `json:"plate"`
	Zone           string `json:"zone"`
	Station        string `json:"station"`
	Unit_type      string `json:"unit_type"`
	Make           string `json:"make"`
	Drivers        int    `json:"drivers"`
	Unit_condition string `json:"unit_condition"`
	Vehicle_serial string `json:"vehicle_serial"`
	Motor_serial   string `json:"motor_serial"`
	Capacity       string `json:"capacity"`
	Details        string `json:"details"`
	Fuel_type      string `json:"fuel_type"`
	Water_capacity string `json:"water_capacity"`
	Observations   string `json:"observations"`
}

type UnitDto struct {
	Id             int    `json:"id"`
	Plate          string `json:"plate"`
	Zone           string `json:"zone"`
	Station        string `json:"station"`
	Unit_type      string `json:"unit_type"`
	Make           string `json:"make"`
	Drivers        int    `json:"drivers"`
	Unit_condition string `json:"unit_condition"`
	Vehicle_serial string `json:"vehicle_serial"`
	Motor_serial   string `json:"motor_serial"`
	Capacity       string `json:"capacity"`
	Details        string `json:"details"`
	Fuel_type      string `json:"fuel_type"`
	Water_capacity string `json:"water_capacity"`
	Observations   string `json:"observations"`
}