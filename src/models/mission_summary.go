package models

import "database/sql"

type MissionSummary struct {
	Id              int64          `db:"id"`
	Alias           sql.NullString `db:"alias"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	Code            sql.NullString `db:"code"`
	NumServices     sql.NullInt64  `db:"num_services"`
	NumFireFighters sql.NullInt64  `db:"num_firefighters"`
	NumVehicles     sql.NullInt64  `db:"num_vehicles"`
	Unharmed        sql.NullInt64  `db:"unharmed"`
	Injured         sql.NullInt64  `db:"injured"`
	Transported     sql.NullInt64  `db:"transported"`
	Deceased        sql.NullInt64  `db:"deceased"`
	OperativeAreas  []sql.NullString `db:"operative_areas"`
}
