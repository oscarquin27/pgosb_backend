package state_entity

import "github.com/jackc/pgx/v5/pgtype"

type State struct {
	Id          pgtype.Int8 `json:"id" db:"state_id"`
	Name    pgtype.Text `json:"name"`
	Coordinates pgtype.Text `json:"coordinates"`
}