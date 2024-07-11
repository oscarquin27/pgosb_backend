package api_models

import (
	"fdms/src/models"
	"fmt"
	"strconv"
	"time"
)

type RoleJson struct {
	Id           string           `json:"id"`
	RoleName     string           `json:"role_name" `
	AccessSchema AccessSchemaJson `json:"access_schema"`
	StRole       int              `json:"st_role"`
	Created_at   string           `json:"created_at"`
	Updated_at   string           `json:"updated_at"`
}

// AccessSchemaJson mirrors AccessSchema but uses string keys
type AccessSchemaJson struct {
	Roles               map[string]string `json:"roles,omitempty"`
	Units               map[string]string `json:"units,omitempty"`
	Users               map[string]string `json:"users,omitempty"`
	Services            map[string]string `json:"services,omitempty"`
	Stations            map[string]string `json:"stations,omitempty"`
	Locations           map[string]string `json:"locations,omitempty"`
	AssistentialCenters map[string]string `json:"assistential_centers,omitempty"`
}

// RoleToDto converts a Role struct to a RoleJson struct.
func ModelToRoleJson(role *models.Role) (*RoleJson, error) {
	accessSchemaDto := AccessSchemaJson{
		Roles:               map[string]string{},
		Units:               map[string]string{},
		Users:               map[string]string{},
		Services:            map[string]string{},
		Stations:            map[string]string{},
		Locations:           map[string]string{},
		AssistentialCenters: map[string]string{},
	}

	// Populate AccessSchemaJson from AccessSchema
	addPermissionsToDto(&accessSchemaDto.Roles, role.AccessSchema.Roles)
	addPermissionsToDto(&accessSchemaDto.Units, role.AccessSchema.Units)
	addPermissionsToDto(&accessSchemaDto.Users, role.AccessSchema.Users)
	addPermissionsToDto(&accessSchemaDto.Services, role.AccessSchema.Services)
	addPermissionsToDto(&accessSchemaDto.Stations, role.AccessSchema.Stations)
	addPermissionsToDto(&accessSchemaDto.Locations, role.AccessSchema.Locations)
	addPermissionsToDto(&accessSchemaDto.AssistentialCenters, role.AccessSchema.AssistentialCenters)

	return &RoleJson{
		Id:           fmt.Sprint(role.ID),
		RoleName:     role.RoleName,
		AccessSchema: accessSchemaDto,
		StRole:       role.StRole,
		Created_at:   role.CreatedAt.Format(time.RFC3339),
		Updated_at:   role.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func addPermissionsToDto(dtoPermissions *map[string]string, permissions models.Permissions) {
	if len(permissions) > 0 {
		*dtoPermissions = make(map[string]string)
		for k, v := range permissions {
			(*dtoPermissions)[k] = fmt.Sprint(v)
		}
	}
}

// addDtoPermissionsToRole add dto permissions to the role if the dto permissions are not empty
func addDtoPermissionsToRole(permissions *models.Permissions, dtoPermissions map[string]string) error {
	if len(dtoPermissions) > 0 {
		*permissions = make(models.Permissions)
		for k, v := range dtoPermissions {
			boolValue, err := strconv.ParseBool(v)
			if err != nil {
				return fmt.Errorf("error parsing access schema value: %w", err)
			}
			(*permissions)[k] = boolValue
		}
	}
	return nil
}

// DtoToRole converts a RoleJson struct to a Role struct.
func (dto *RoleJson) ToModel() (*models.Role, error) {

	accessSchema := models.AccessSchema{
		Roles:               models.Permissions{},
		Units:               models.Permissions{},
		Users:               models.Permissions{},
		Services:            models.Permissions{},
		Stations:            models.Permissions{},
		Locations:           models.Permissions{},
		AssistentialCenters: models.Permissions{},
	}

	addDtoPermissionsToRole(&accessSchema.Roles, dto.AccessSchema.Roles)
	addDtoPermissionsToRole(&accessSchema.Units, dto.AccessSchema.Units)
	addDtoPermissionsToRole(&accessSchema.Users, dto.AccessSchema.Users)
	addDtoPermissionsToRole(&accessSchema.Services, dto.AccessSchema.Services)
	addDtoPermissionsToRole(&accessSchema.Stations, dto.AccessSchema.Stations)
	addDtoPermissionsToRole(&accessSchema.Locations, dto.AccessSchema.Locations)
	addDtoPermissionsToRole(&accessSchema.AssistentialCenters, dto.AccessSchema.AssistentialCenters)

	id, err := strconv.ParseInt(dto.Id, 10, 64)
	if err != nil {
		return &models.Role{}, fmt.Errorf("error converting id to integer: %w", err)
	}

	// stRole, err := strconv.Atoi(dto.StRole)
	// if err != nil {
	// 	return Role{}, fmt.Errorf("error converting st_role to integer: %w", err)
	// }

	createdAt, err := time.Parse(time.RFC3339, dto.Created_at)
	if err != nil {
		return &models.Role{}, fmt.Errorf("error parsing created_at: %w", err)
	}

	updatedAt, err := time.Parse(time.RFC3339, dto.Updated_at)
	if err != nil {
		return &models.Role{}, fmt.Errorf("error parsing updated_at: %w", err)
	}

	return &models.Role{
		ID:           id,
		RoleName:     dto.RoleName,
		StRole:       dto.StRole,
		AccessSchema: accessSchema,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}
