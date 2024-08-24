package models

type MissionFirefighter struct {
	Id            int64   `db:"id"`
	MissionId     int64   `db:"mission_id"`
	ServiceId     int64   `db:"service_id"`
	FireFighterId int64   `db:"firefigther_id"`
	ServiceRol    *string `db:"service_role"`
}
