package models

import "github.com/jackc/pgx/v5"

type MissionFirefighter struct {
	Id int64 `db:"id"`
	//ServiceId  int64   `db:"service_id"`
	UserId     int64   `db:"user_id"`
	ServiceRol *string `db:"service_role"`
	MissionId  int64   `db:"mission_id"`
}

func (s *MissionFirefighter) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id": s.Id,
		//"service_id":   s.ServiceId,
		"user_id":      s.UserId,
		"service_role": s.ServiceRol,
		"mission_id":   s.MissionId,
	}
}

func (s *MissionFirefighter) SetId(id int64) {
	s.Id = id
}
