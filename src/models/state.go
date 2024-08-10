package models

import "github.com/jackc/pgx/v5"

type State struct {
	Id          int64   `db:"id"`
	Name        *string `db:"name"`
	Capital     *string `db:"capital"`
	Coordinates *string `db:"coordinates"`
}

func (s *State) SetId(id int64) {
	s.Id = id
}

func (s *State) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":          s.Id,
		"name":        s.Name,
		"coordinates": s.Coordinates,
		"capital":     s.Capital,
	}
}
