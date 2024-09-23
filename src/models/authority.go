package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type Authority struct {
	Id           int64          `db:"id"`
	Name         sql.NullString `db:"name"`
	Abbreviation sql.NullString `db:"abbreviation"`
	Government   sql.NullString `db:"government"`
}

func (s *Authority) SetId(id int64) {
	s.Id = id
}

func (s *Authority) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id

	if s.Name.Valid {
		args["name"] = s.Name.String
	}
	if s.Abbreviation.Valid {
		args["abbreviation"] = s.Abbreviation.String
	}
	if s.Government.Valid {
		args["government"] = s.Government.String
	}

	return args
}
