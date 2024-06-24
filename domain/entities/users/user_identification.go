package user_entity

import "github.com/jackc/pgx/v5/pgtype"

type UserIdentification struct {
	Id      pgtype.Int8   `json:"id"`
	Id_role pgtype.Int8   `json:"id_role" binding:"required"`
}