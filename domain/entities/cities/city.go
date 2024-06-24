package city_entity

import "github.com/jackc/pgx/v5/pgtype"

type City struct {
	Id        pgtype.Int8 `json:"id" db:"city_id"`
	State_Id  pgtype.Int8 `json:"state_id"`
	Name      pgtype.Text `json:"name"`
	Area_Code pgtype.Text `json:"area_code"`
	Zip_Code pgtype.Text `json:"zip_code"`
	Coordinates pgtype.Text `json:"Coordinates"`
}