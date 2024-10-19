package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionFirefighter struct {
	Id         int64        `db:"id"`
	ServiceId  int64        `db:"service_id"`
	MissionId  int64        `db:"mission_id"`
	UserId     int64        `db:"user_id"`
	ServiceRol *string      `db:"service_role"`
	CreatedAt  sql.NullTime `db:"created_at"`
}

func (s *MissionFirefighter) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":           s.Id,
		"service_id":   s.ServiceId,
		"user_id":      s.UserId,
		"service_role": s.ServiceRol,
		"mission_id":   s.MissionId,
		"created_at":   s.CreatedAt,
	}
}

func (s *MissionFirefighter) SetId(id int64) {
	s.Id = id
}

type MissionFirefighterUser struct {
	Id           int64   `db:"id"`
	UserId       int64   `db:"user_id"`
	Name         string  `db:"name"`
	User_name    *string `db:"user_name"`
	Rank         string  `db:"rank"`
	PersonalCode string  `db:"personal_code"`
	Legal_id     string  `db:"legal_id"`
	MissionId    string  `db:"mission_id"`
}

func (s *MissionFirefighterUser) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":            s.Id,
		"user_id":       s.UserId,
		"name":          s.Name,
		"user_name":     s.User_name,
		"rank":          s.Rank,
		"personal_code": s.PersonalCode,
		"legal_id":      s.Legal_id,
		"mission_id":    s.MissionId,
	}
}

func (s *MissionFirefighterUser) SetId(id int64) {
	s.Id = id
}
