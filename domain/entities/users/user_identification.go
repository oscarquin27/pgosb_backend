package user_entity

import "github.com/jackc/pgx/v5/pgtype"

type UserIdentification struct {
	Id      pgtype.Text   `json:"id"`
	Id_role pgtype.Text   `json:"id_role" binding:"required"`
}