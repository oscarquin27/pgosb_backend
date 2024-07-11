package models

import (
	"errors"
	"time"
)

var (
	ErrorRoleNotFound   = errors.New("rol no encontrado")
	ErrorRoleNotCreated = errors.New("rol no creado")
	ErrorRoleNotUpdated = errors.New("el rol no pudo ser actualizado")
	ErrorRoleNotDeleted = errors.New("el rol no pudo ser eliminado")
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
	ID           int64        `json:"id"`
	RoleName     string       `json:"role_name"`
	StRole       int          `json:"st_role"`
	AccessSchema AccessSchema `json:"access_schema"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	UpdatedAt    time.Time    `json:"updated_at,omitempty"`
}
