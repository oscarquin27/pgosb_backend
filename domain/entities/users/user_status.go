package user_entity

import (
	"time"
)

type UserStatus struct {
	Status_user string `json:"status_user"`
	Last_connection time.Time `json:"last_connection"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Ip string `json:"ip"`
}

type UserStatusDto struct {
	Status_user string `json:"status_user"`
	Last_connection string `json:"last_connection"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Ip string `json:"ip"`
}