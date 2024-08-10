package models

import (
	"time"
)

type Permissions map[string]bool

type AccessSchema struct {
	Roles               Permissions `json:"roles"`
	Units               Permissions `json:"units"`
	Users               Permissions `json:"users"`
	Services            Permissions `json:"services"`
	Stations            Permissions `json:"stations"`
	Locations           Permissions `json:"locations"`
	AssistentialCenters Permissions `json:"assistential_centers"`
}

type Role struct {
	ID           int64        `db:"id"`
	RoleName     string       `db:"role_name"`
	StRole       bool         `db:"st_role"`
	AccessSchema AccessSchema `db:"access_schema"`
	CreatedAt    time.Time    `db:"created_at,omitempty"`
	UpdatedAt    time.Time    `db:"updated_at,omitempty"`
}
