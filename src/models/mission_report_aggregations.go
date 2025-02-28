package models

import "database/sql"

// AntaresStationAggregation represents the result of the antares_station query
// select antares_id, antares_name, station_id, station_name, station_abbreviation, count(1) from mission_report_summary
type AntaresStationAggregation struct {
	AntaresId           sql.NullInt64  `db:"antares_id"`
	AntaresName         sql.NullString `db:"antares_name"`
	StationId           sql.NullInt64  `db:"station_id"`
	StationName         sql.NullString `db:"station_name"`
	StationAbbreviation sql.NullString `db:"station_abbreviation"`
	Count               int64          `db:"count"`

	Unharmed    sql.NullInt64 `db:"unharmed"`
	Injured     sql.NullInt64 `db:"injured"`
	Transported sql.NullInt64 `db:"transported"`
	Deceased    sql.NullInt64 `db:"deceased"`
}

// AntaresAggregation represents the result of the antares query
// select antares_id, antares_name, count(1) from mission_report_summary
type AntaresAggregation struct {
	AntaresId   sql.NullInt64  `db:"antares_id"`
	AntaresName sql.NullString `db:"antares_name"`
	Count       int64          `db:"count"`
}

// StationAggregation represents the result of the station query
// select station_id, station_abbreviation, station_name, count(1) from mission_report_summary
type StationAggregation struct {
	StationId           sql.NullInt64  `db:"station_id"`
	StationName         sql.NullString `db:"station_name"`
	StationAbbreviation sql.NullString `db:"station_abbreviation"`
	Count               int64          `db:"count"`
}

// AntaresTypeAggregation represents the result of the antares_type query
// select antares_type, count(1) from mission_report_summary
type AntaresTypeAggregation struct {
	AntaresType sql.NullString `db:"antares_type"`
	Count       int64          `db:"count"`
}
