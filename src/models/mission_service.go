package models

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionServiceNotFound   = errors.New("servicio no encontrado")
	ErrorMissionServiceNotCreated = errors.New("servicio no creado")
	ErrorMissionServiceNotUpdated = errors.New("el servicio no pudo ser actualizado")
	ErrorMissionServiceNotDeleted = errors.New("el servicio no pudo ser eliminado")
)

type MissionService struct {
	Id        pgtype.Int4 `json:"id" db:"id"`
	MissionId pgtype.Int2 `json:"mission_id" db:"mission_id"`
	AntaresId pgtype.Int2 `json:"antares_id,omitempty" db:"antares_id"`
	//Units          []pgtype.Int2 `json:"units,omitempty" db:"units"`
	//Bombers        []pgtype.Int2 `json:"bombers,omitempty" db:"Bombers"`
	//OperativeAreas []pgtype.Text `json:"operative_areas,omitempty" db:"operative_areas"`

	Summary     pgtype.Text `json:"summary,omitempty" db:"summary"`
	Description pgtype.Text `json:"description,omitempty" db:"description"`

	//Unharmed    *int64 `json:"unharmed" db:"unharmed"`
	//Injured     *int64 `json:"injured" db:"injured"`
	//Transported *int64 `json:"transported" db:"transported"`
	//Deceased    *int64 `json:"deceased" db:"deceased"`

	StationId         *int64 `json:"station_id" db:"station_id"`
	LocationId        *int64 `json:"location_id" db:"location_id"`
	LocationDestinyId *int64 `json:"location_destiny_id" db:"location_destiny_id"`

	ServiceDate       sql.NullTime `json:"service_date" db:"service_date"`
	ManualServiceDate sql.NullTime `json:"manual_service_date" db:"manual_service_date"`
	IsImportant       bool         `json:"is_important" db:"is_important"`

	SendingUserId   *int64 `json:"sending_user_id" db:"sending_user_id"`
	ReceivingUserId *int64 `json:"receiving_user_id" db:"receiving_user_id"`

	HealthCareCenterId *int64 `json:"center_id" db:"center_id"`

	Level         sql.NullString `json:"level" db:"level"`
	PeaceQuadrant sql.NullString `json:"peace_quadrant" db:"peace_quadrant"`

	CanceledReason sql.NullString `json:"cancel_reason" db:"cancel_reason"`
	PendingForData sql.NullBool   `json:"pending_for_data" db:"pending_for_data"`
}

type RelevantServices struct {
	Id                    int64                     `db:"id"`
	RegionAreaDescription *string                   `db:"region_area"`
	MissionCode           *string                   `db:"mission_code"`
	AntaresId             *int                      `db:"antares_id"`
	AntaresType           *string                   `db:"antares_type"`
	AntaresDescription    *string                   `db:"antares_description"`
	ServiceId             *int                      `db:"service_id"`
	OprativeAreaName      []string                  `db:"operative_area_name"`
	ServiceDescription    *string                   `db:"service_description"`
	ServiceDate           *string                   `db:"service_date"`
	Units                 []string                  `db:"units"`
	Firefighters          []RelevantFirefighters    `db:"firefighters"`
	People                []RelevantPeople          `db:"people"`
	Infrastructures       []RelevantInfrastructure  `db:"infrastructures"`
	Vehicles              []RelevantVehicle         `db:"vehicles"`
	ServiceLocations      []RelevantServiceLocation `db:"service_locations"`
	ServiceStations       []RelevantServiceStation  `db:"service_stations"`
	Centers               []RelevantCenter          `db:"centers"`
	Unharmed              *int64                    `db:"unharmed"`
	Injured               *int64                    `db:"injured"`
	Transported           *int64                    `db:"transported"`
	Deceased              *int64                    `db:"deceased"`
	IsImportant           bool                      `db:"is_important"`
}

type RelevantFirefighters struct {
	Rank     *string `json:"rank"`
	Name     *string `json:"name"`
	Document *string `json:"document"`
	Role     *string `json:"role"`
	Team     *string `json:"team"`
}

type RelevantPeople struct {
	Condition       *string `json:"condition"`
	Name            *string `json:"name"`
	Gender          *string `json:"gender"`
	Age             *string `json:"age"`
	Document        *string `json:"document"`
	Phone           *string `json:"phone"`
	PersonCondition *string `json:"person_condition"`
	Unit            *string `json:"unit"` // May need adjustment based on vehicle data structure
	Address         *string `json:"address"`
	Building        *string `json:"building"`
	Vehicle         *string `json:"vehicle"`
}

type RelevantInfrastructure struct {
	Type       *string `json:"type"`
	Floor      *string `json:"floor"`
	Occupation *string `json:"occupation"`
	Levels     *string `json:"levels"` // Assuming levels is an integer
}

type RelevantVehicle struct {
	Plate  *string `json:"plate"`
	Make   *string `json:"make"`
	Model  *string `json:"model"`
	Year   *string `json:"year"` // Assuming year is an integer
	Color  *string `json:"color"`
	Type   *string `json:"type"`
	Serial *string `json:"serial"`
}

type RelevantServiceLocation struct {
	State        *string `json:"state"`
	Municipality *string `json:"municipality"`
	Parish       *string `json:"parish"`
	Sector       *string `json:"sector"`
	Urb          *string `json:"urb"`
	Address      *string `json:"address"`
}

type RelevantServiceStation struct {
	Name         *string `json:"name"`
	Abbreviation *string `json:"abbreviation"`
	State        *string `json:"state"`
	Municipality *string `json:"municipality"`
	Parish       *string `json:"parish"`
	Sector       *string `json:"sector"`
	Urb          *string `json:"urb"`
}

type RelevantCenter struct {
	Name         *string `json:"name"`
	Abbreviation *string `json:"abbreviation"`
	State        *string `json:"state"`
	Municipality *string `json:"municipality"`
	Parish       *string `json:"parish"`
	Sector       *string `json:"sector"`
	Urb          *string `json:"urb"`
}
