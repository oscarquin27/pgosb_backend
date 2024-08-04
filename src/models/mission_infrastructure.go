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
	Id              pgtype.Int4 `json:"id" binding:"required"`
	ServiceId       pgtype.Int4 `json:"service_id"`
	BuildType       pgtype.Text `json:"build_type"`
	BuildOccupation pgtype.Text `json:"build_occupation"`
	BuildArea       pgtype.Text `json:"build_area"`
	BuildAccess     pgtype.Text `json:"build_access"`
	Levels          pgtype.Int2 `json:"levels"`
	Walls           pgtype.Text `json:"build_wall"`
	People          pgtype.Int2 `json:"people"`
	GoodsType       pgtype.Text `json:"goods_type"`
	BuildRoof       pgtype.Text `json:"build_roof"`
	BuildFloor      pgtype.Text `json:"build_floor"`
	BuildRoomType   pgtype.Text `json:"build_room_type"`
	Observations    pgtype.Text `json:"observations"`
}
