package models

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type MissionAuthorityPerson struct {
	Id                   int64          `db:"id"`
	MissionId            int64          `db:"mission_id"`
	AuthorityId          int64          `db:"authority_id"`
	Name                 sql.NullString `db:"name"`
	LastName             sql.NullString `db:"last_name"`
	LegalId              sql.NullString `db:"legal_id"`
	IdentificationNumber sql.NullString `db:"identification_number"`
	Phone                sql.NullString `db:"phone"`
	Gender               sql.NullString `db:"gender"`
	Observations         sql.NullString `db:"observations"`
	CreatedAt            sql.NullTime   `db:"created_at"`
}

func (s *MissionAuthorityPerson) SetId(id int64) {
	s.Id = id
}

func (s *MissionAuthorityPerson) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = s.Id
	args["mission_id"] = s.MissionId

	if s.Name.Valid {
		args["name"] = s.Name.String
	}

	if s.LastName.Valid {
		args["last_name"] = s.LastName.String
	}

	if s.LegalId.Valid {
		args["legal_id"] = s.LegalId.String
	}

	if s.IdentificationNumber.Valid {
		args["identification_number"] = s.IdentificationNumber.String
	}

	if s.Phone.Valid {
		args["phone"] = s.Phone.String
	}

	if s.Gender.Valid {
		args["gender"] = s.Gender.String
	}

	if s.Observations.Valid {
		args["observations"] = s.Observations.String
	}

	args["authority_id"] = s.AuthorityId

	return args

}
