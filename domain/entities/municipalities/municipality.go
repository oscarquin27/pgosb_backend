package municipality_entity

import "github.com/jackc/pgx/v5/pgtype"

type Municipality struct {
	Id pgtype.Int8 `json:"id"`
	State_Id  pgtype.Int8 `json:"state_id"`
	Name  pgtype.Text `json:"name"`
	Coordinates pgtype.Text `json:"coordinates"`
}