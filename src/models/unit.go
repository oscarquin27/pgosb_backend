package models

type Unit struct {
	Id             int64    `db:"id"`
	Plate          *string  `db:"plate"`
	Station        *string  `db:"station"`
	Unit_type      *string  `db:"unit_type"`
	Make           *string  `db:"make"`
	Unit_condition *string  `db:"unit_condition"`
	Vehicle_serial *string  `db:"vehicle_serial"`
	Motor_serial   *string  `db:"motor_serial"`
	Capacity       *string  `db:"capacity"`
	Details        []string `db:"details"`
	Fuel_type      *string  `db:"fuel_type"`
	Water_capacity *string  `db:"water_capacity"`
	Observations   *string  `db:"observations"`
	Hurt_capacity  *int64   `db:"hurt_capacity"`
	Doors          *int64   `db:"doors"`
	Performance    *string  `db:"performance"`
	Load_capacity  *int64   `db:"load_capacity"`
	Model          *string  `db:"model"`
	Alias          *string  `db:"alias"`
	Color          *string  `db:"color"`
	Year           *string  `db:"year"`
	Purpose        *string  `db:"purpose"`
	Init_kilometer *int64   `db:"init_kilometer"`
}
