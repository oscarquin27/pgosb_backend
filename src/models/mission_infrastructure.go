package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorMissionInfrastructureNotFound   = errors.New("infraestructura no encontrada")
	ErrorMissionInfrastructureNotCreated = errors.New("infraestructura no creada")
	ErrorMissionInfrastructureNotUpdated = errors.New("la infraestructura no pudo ser actualizada")
	ErrorMissionInfrastructureNotDeleted = errors.New("la infraestructura no pudo ser eliminada")
)

type MissionInfrastructure struct {
	Id              pgtype.Int4 `json:"id" db:"id"`
	ServiceId       pgtype.Int4 `json:"service_id" db:"service_id"`
	BuildType       pgtype.Text `json:"build_type" db:"build_type"`
	BuildOccupation pgtype.Text `json:"build_occupation" db:"build_occupation"`
	BuildArea       pgtype.Text `json:"build_area" db:"build_area"`
	BuildAccess     pgtype.Text `json:"build_access" db:"build_access"`
	Levels          pgtype.Int2 `json:"levels" db:"levels"`
	Walls           pgtype.Text `json:"build_wall" db:"build_wall"`
	People          pgtype.Int2 `json:"people" db:"people"`
	GoodsType       pgtype.Text `json:"goods_type" db:"goods_type"`
	BuildRoof       pgtype.Text `json:"build_roof" db:"build_roof"`
	BuildFloor      pgtype.Text `json:"build_floor" db:"build_floor"`
	BuildRoomType   pgtype.Text `json:"build_room_type" db:"build_room_type"`
	Observations    pgtype.Text `json:"observations" db:"observations"`
}
