package models

type State struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Area_Code   string `db:"area_code"`
	Zip_Code    string `db:"zip_code"`
	Coordinates string `db:"Coordinates"`
}
