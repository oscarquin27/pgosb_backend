package role_entity

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Role struct {
	ID           int            `db:"id"`
	RoleName     string         `db:"role_name"`
	StRole       int            `db:"st_role"`
	AccessSchema map[string]any `db:"access_schema"` // Use 'any' for flexibility
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}

type RoleDto struct {
	Id           string `json:"id"`
	RoleName     string `json:"role_name" `
	AccessSchema string `json:"access_schema"`
	StRole       string `json:"st_role"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}

func RoleToDto(role Role) (RoleDto, error) {
	// Convert AccessSchema to JSON string
	accessSchemaJSON, err := json.Marshal(role.AccessSchema)
	if err != nil {
		return RoleDto{}, fmt.Errorf("error marshalling access schema: %w", err)
	}

	return RoleDto{
		Id:           fmt.Sprint(role.ID), // Convert int to string
		RoleName:     role.RoleName,
		AccessSchema: string(accessSchemaJSON),
		StRole:       fmt.Sprint(role.StRole), // Convert int to string
		Created_at:   role.CreatedAt.Format(time.RFC3339),
		Updated_at:   role.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// DtoToRole converts a RoleDto struct to a Role struct.
func DtoToRole(dto RoleDto) (Role, error) {
	var accessSchema map[string]any

	// Parse AccessSchema from JSON string
	err := json.Unmarshal([]byte(dto.AccessSchema), &accessSchema)
	if err != nil {
		return Role{}, fmt.Errorf("error unmarshalling access schema: %w", err)
	}

	id, err := strconv.Atoi(dto.Id)
	if err != nil {
		return Role{}, fmt.Errorf("error converting id to integer: %w", err)
	}

	stRole, err := strconv.Atoi(dto.StRole)
	if err != nil {
		return Role{}, fmt.Errorf("error converting st_role to integer: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, dto.Created_at)
	if err != nil {
		return Role{}, fmt.Errorf("error parsing created_at: %w", err)
	}

	updatedAt, err := time.Parse(time.RFC3339, dto.Updated_at)
	if err != nil {
		return Role{}, fmt.Errorf("error parsing updated_at: %w", err)
	}

	return Role{
		ID:           id,
		RoleName:     dto.RoleName,
		StRole:       stRole,
		AccessSchema: accessSchema,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}
