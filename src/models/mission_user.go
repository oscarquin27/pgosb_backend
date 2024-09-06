package models

import "database/sql"

type MissionUserService struct {
	Id           int64          `db:"id"`
	ServiceId    int64          `db:"service_id"`
	UserId       sql.NullInt64  `db:"user_id"`
	Name         sql.NullString `db:"name"`
	UserName     sql.NullString `db:"user_name"`
	Rank         sql.NullString `db:"rank"`
	PersonalCode sql.NullString `db:"personal_code"`
	LegalId      sql.NullString `db:"legal_id"`
	ServiceRole  sql.NullString `db:"service_role"`
}
