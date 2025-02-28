package models

import "database/sql"

type MissionReportSummary struct {
	MissionId           int64            `db:"mission_id"`
	AntaresId           sql.NullInt64    `db:"antares_id"`
	Id                  int64            `db:"id"`
	CreatedAt           sql.NullTime     `db:"created_at"`
	Code                sql.NullString   `db:"code"`
	Alias               sql.NullString   `db:"alias"`
	OperativeAreas      []sql.NullString `db:"operative_areas"`
	Summary             sql.NullString   `db:"summary"`
	Description         sql.NullString   `db:"description"`
	Unharmed            sql.NullInt64    `db:"unharmed"`
	Injured             sql.NullInt64    `db:"injured"`
	Transported         sql.NullInt64    `db:"transported"`
	Deceased            sql.NullInt64    `db:"deceased"`
	StationId           sql.NullInt64    `db:"station_id"`
	LocationId          sql.NullInt64    `db:"location_id"`
	ManualMissionDate   sql.NullTime     `db:"manual_mission_date"`
	IsImportant         bool             `db:"is_important"`
	CenterId            sql.NullInt64    `db:"center_id"`
	SendingUserId       sql.NullInt64    `db:"sending_user_id"`
	Level               sql.NullString   `db:"level"`
	PeaceQuadrant       sql.NullString   `db:"peace_quadrant"`
	LocationDestinyId   sql.NullInt64    `db:"location_destiny_id"`
	CancelReason        sql.NullString   `db:"cancel_reason"`
	PendingForData      bool             `db:"pending_for_data"`
	ReceivingUserId     sql.NullInt64    `db:"receiving_user_id"`
	AntaresType         sql.NullString   `db:"antares_type"`
	AntaresName         sql.NullString   `db:"antares_name"`
	StationName         sql.NullString   `db:"station_name"`
	StationAbbreviation sql.NullString   `db:"station_abbreviation"`
}
