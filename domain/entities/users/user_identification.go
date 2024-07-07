package user_entity

import "github.com/jackc/pgx/v5/pgtype"

type UserIdentification struct {
	Id      int64       `json:"id"`
	Id_role pgtype.Int4 `json:"id_role"`
}

type UserIdentificationDto struct {
	Id      string `json:"id"`
	Id_role string `json:"id_role"`
}
