package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionPersonNotFound   = errors.New("persona no encontrado")
	ErrorMissionPersonNotCreated = errors.New("persona no creado")
	ErrorMissionPersonNotUpdated = errors.New("la persona no pudo ser actualizada")
	ErrorMissionPersonNotDeleted = errors.New("la persona no pudo ser eliminada")
)

type MissionPerson struct {
	Id               pgtype.Int4 `json:"id" db:"id"`
	MissionId        pgtype.Int4 `json:"mission_id" db:"mission_id"`
	UnitId           pgtype.Int4 `json:"unit_id" db:"unit_id"`
	InfrastructureId pgtype.Int4 `json:"infrastructure_id" db:"infrastructure_id"`
	VehicleId        pgtype.Int4 `json:"vehicle_id" db:"vehicle_id"`
	FirstName        pgtype.Text `json:"first_name" db:"first_name"`
	LastName         pgtype.Text `json:"last_name" db:"last_name"`
	Age              pgtype.Int2 `json:"age" db:"age"`
	Gender           pgtype.Text `json:"gender" db:"gender"`
	LegalId          pgtype.Text `json:"legal_id" db:"legal_id"`
	Phone            pgtype.Text `json:"phone" db:"phone"`
	Employment       pgtype.Text `json:"employment" db:"employment"`
	State            pgtype.Text `json:"state" db:"state"`
	Municipality     pgtype.Text `json:"municipality" db:"municipality"`
	Parish           pgtype.Text `json:"parish" db:"parish"`
	Address          pgtype.Text `json:"address" db:"address"`
	Pathology        pgtype.Text `json:"pathology" db:"pathology"`
	Observations     pgtype.Text `json:"observations" db:"observations"`
	Condition        pgtype.Text `json:"condition" db:"condition"`

	//Temporal
	ServiceId pgtype.Int4 `json:"service_id" db:"service_id"`
}
