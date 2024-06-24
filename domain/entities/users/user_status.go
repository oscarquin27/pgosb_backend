package user_entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UserStatus struct {
	Status_user pgtype.Int8 `json:"status_user"`
	Last_connection pgtype.Timestamp `json:"last_connection"`
	Created_at pgtype.Timestamp `json:"created_at"`
	Updated_at pgtype.Timestamp `json:"updated_at"`
	Ip pgtype.Text `json:"ip"`
}
