package models

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrorMissionNotFound   = errors.New("misión no encontrada")
	ErrorMissionNotCreated = errors.New("misión no creada")
	ErrorMissionNotUpdated = errors.New("la misión pudo ser actualizado")
	ErrorMissionNotDeleted = errors.New("la misión no pudo ser eliminada")
)

type Mission struct {
	Id        int64          `db:"id"`
	CreatedAt sql.NullTime   `db:"created_at,omitempty"`
	Code      sql.NullString `db:"code"`
	Alias     sql.NullString `db:"alias"`

	OperativeAreas []sql.NullString `json:"operative_areas,omitempty" db:"operative_areas"`
	Summary        sql.NullString   `json:"summary,omitempty" db:"summary"`
	Description    sql.NullString   `json:"description,omitempty" db:"description"`

	Unharmed    sql.NullInt64 `json:"unharmed" db:"unharmed"`
	Injured     sql.NullInt64 `json:"injured" db:"injured"`
	Transported sql.NullInt64 `json:"transported" db:"transported"`
	Deceased    sql.NullInt64 `json:"deceased" db:"deceased"`

	StationId  sql.NullInt64 `json:"station_id" db:"station_id"`
	LocationId sql.NullInt64 `json:"location_id" db:"location_id"`

	//ServiceDate       sql.NullTime `json:"service_date" db:"service_date"`
	ManualMissionDate sql.NullTime `json:"manual_mission_date" db:"manual_mission_date"`
	IsImportant       sql.NullBool `json:"is_important" db:"is_important"`

	HealthCareCenterId sql.NullInt64 `json:"center_id" db:"center_id"`
	SendingUserId      sql.NullInt64 `json:"sending_user_id" db:"sending_user_id"`
	ReceivingUserId    sql.NullInt64 `json:"receiving_user_id" db:"receiving_user_id"`

	Level         sql.NullString `json:"level" db:"level"`
	PeaceQuadrant sql.NullString `json:"peace_quadrant" db:"peace_quadrant"`

	LocationDestinyId sql.NullInt64 `json:"location_destiny_id" db:"location_destiny_id"`

	PendingForData sql.NullBool   `json:"pending_for_data" db:"pending_for_data"`
	CanceledReason sql.NullString `json:"cancel_reason" db:"cancel_reason"`
}

func (m *Mission) SetId(id int64) {
	m.Id = id
}

func (m *Mission) GetNameArgs() pgx.NamedArgs {

	args := make(pgx.NamedArgs)

	args["id"] = m.Id

	if m.CreatedAt.Valid {
		args["created_at"] = m.CreatedAt.Time
	}
	if m.Code.Valid {
		args["code"] = m.Code.String
	}
	if m.Alias.Valid {
		args["alias"] = m.Alias.String
	}
	if len(m.OperativeAreas) > 0 {
		args["operative_areas"] = m.OperativeAreas
	}
	if m.Summary.Valid {
		args["summary"] = m.Summary.String
	}
	if m.Description.Valid {
		args["description"] = m.Description.String
	}
	if m.Unharmed.Valid {
		args["unharmed"] = m.Unharmed.Int64
	}
	if m.Injured.Valid {
		args["injured"] = m.Injured.Int64
	}
	if m.Transported.Valid {
		args["transported"] = m.Transported.Int64
	}
	if m.Deceased.Valid {
		args["deceased"] = m.Deceased.Int64
	}
	if m.StationId.Valid {
		args["station_id"] = m.StationId.Int64
	}
	if m.LocationId.Valid {
		args["location_id"] = m.LocationId.Int64
	}
	//if m.ServiceDate.Valid {
	//	args["service_date"] = m.ServiceDate.Time
	//}
	if m.ManualMissionDate.Valid {
		args["manual_mission_date"] = m.ManualMissionDate.Time
	}
	if m.IsImportant.Valid {
		args["is_important"] = m.IsImportant.Bool
	}
	if m.HealthCareCenterId.Valid {
		args["health_care_center_id"] = m.HealthCareCenterId.Int64
	}
	if m.SendingUserId.Valid {
		args["sending_user_id"] = m.SendingUserId.Int64
	}
	if m.ReceivingUserId.Valid {
		args["receiving_user_id"] = m.ReceivingUserId.Int64
	}
	if m.Level.Valid {
		args["level"] = m.Level.String
	}
	if m.PeaceQuadrant.Valid {
		args["peace_quadrant"] = m.PeaceQuadrant.String
	}
	if m.LocationDestinyId.Valid {
		args["location_destiny_id"] = m.LocationDestinyId.Int64
	}

	if m.PendingForData.Valid {
		args["pending_for_data"] = m.PendingForData.Bool
	}
	if m.CanceledReason.Valid {
		args["cancel_reason"] = m.CanceledReason.String
	}

	return args
}
