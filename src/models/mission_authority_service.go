package models

import "github.com/jackc/pgx/v5"

type MissionAuthorityService struct {
	Id          int64 `db:"id"`
	MissionId   int64 `db:"mission_id"`
	ServiceId   int64 `db:"service_id"`
	AuthorityId int64 `db:"authority_id"`
}

func (s *MissionAuthorityService) SetId(id int64) {
	s.Id = id
}

func (s *MissionAuthorityService) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id
	args["mission_id"] = s.MissionId
	args["service_id"] = s.ServiceId
	args["authority_id"] = s.AuthorityId

	return args

}
