package models

import "github.com/jackc/pgx/v5"

type OperativeRegions struct {
	Id           int64   `db:"id"`
	Description  *string `db:"description"`
	Abbreviation *string `db:"abbreviation"`
	Phone        *string `db:"phone"`
	Coverage     *string `db:"coverage"`
}

func (s *OperativeRegions) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":           s.Id,
		"description":  s.Description,
		"abbreviation": s.Abbreviation,
		"phone":        s.Phone,
		"coverage":     s.Coverage,
	}
}

func (s *OperativeRegions) SetId(id int64) {
	s.Id = id
}