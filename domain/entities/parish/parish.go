package parish_entity

import "github.com/jackc/pgx/v5/pgtype"
 
type Parish struct {
	Id              pgtype.Int8 `json:"id" db:"parish_id"`
	State_Id        pgtype.Int8 `json:"state_id"`
	Municipality_Id pgtype.Int8 `json:"municipality_id"`
	Name            pgtype.Text `json:"name"`
	Coordinates     pgtype.Text `json:"coordinates"`
}