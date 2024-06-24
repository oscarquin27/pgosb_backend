package station_entity

import "github.com/jackc/pgx/v5/pgtype"

type Station struct {
	Id          pgtype.Int8 `json:"id" db:"station_id"`
	Municipality_id pgtype.Text `json:"municipality_id"`
	Name    pgtype.Text `json:"name"`
	Coordinates pgtype.Text `json:"coordinates"`
}