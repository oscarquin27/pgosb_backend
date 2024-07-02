package station_entity

import "github.com/jackc/pgx/v5/pgtype"

type Station struct {
	Id	pgtype.Int4 `json:"id" db:"station_id"`
	Municipality_id pgtype.Int4 `json:"municipality_id"`
	Name    pgtype.Text `json:"name"`
	Coordinates pgtype.Text `json:"coordinates"`
	Description pgtype.Text `json:"description"`
	Code pgtype.Text `json:"code"`
	Abbreviation pgtype.Text `json:"abbreviation"`
	Phones []Phones `json:"phones"`
	State_id pgtype.Int4 `json:"state_id"`
	Parish_id pgtype.Int4 `json:"parish_id"`
	Sector pgtype.Text `json:"sector"`
	Community pgtype.Text `json:"community"`
	Street pgtype.Text `json:"street"`
	Address pgtype.Text `json:"address"`
}

type StationDto struct {
	Id	string `json:"id"`
	Municipality_id string `json:"municipality_id"`
	Name string `json:"name"`
	Coordinates string `json:"coordinates"`
	Description string `json:"description"`
	Code string `json:"code"`
	Abbreviation string `json:"abbreviation"`
	Phones []Phones `json:"phones"`
	State_id string `json:"state_id"`
	Parish_id string `json:"parish_id"`
	Sector string `json:"sector"`
	Community string `json:"community"`
	Street string `json:"street"`
	Address string `json:"address"`
}

type Phones struct {
	AreaCode string `json:"area_code"`
	Number string `json:"number"`
}