package role_entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Role struct {
	Id            pgtype.Int8          `json:"id"`
	Role_name     pgtype.Text          `json:"role_name" binding:"required"`
	Access_schema RoleAccessSchema 	   `json:"access_schema" binding:"required"`
	St_role       pgtype.Int8          `json:"st_role"`
	Created_at    pgtype.Timestamp     `json:"created_at"`
	Updated_at    pgtype.Timestamp     `json:"updated_at"`
}