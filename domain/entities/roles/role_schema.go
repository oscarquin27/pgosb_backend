package role_entity

import "github.com/jackc/pgx/v5/pgtype"

type RoleAccessSchema struct {
	Misiones       struct{ DefaultActions } `json:"misiones" binding:"required"`
	Administracion struct{ DefaultActions } `json:"administracion" binding:"required"`
	Auditoria      struct{ DefaultActions } `json:"auditoria" binding:"required"`
	Reporteria     struct{ DefaultActions } `json:"reporteria" binding:"required"`
}

type DefaultActions struct {
	Add    pgtype.Bool `json:"add" binding:"required"`
	Update pgtype.Bool `json:"update" binding:"required"`
	Delete pgtype.Bool `json:"delete" binding:"required"`
	Print  pgtype.Bool `json:"print" binding:"required"`
	Export pgtype.Bool `json:"export" binding:"required"`
}