package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type RoleJson struct {
	Id           string           `json:"id"`
	RoleName     string           `json:"role_name" `
	AccessSchema AccessSchemaJson `json:"access_schema"`
	StRole       bool             `json:"st_role"`
	Created_at   string           `json:"created_at"`
	Updated_at   string           `json:"updated_at"`
}

// AccessSchemaJson mirrors AccessSchema but uses string keys
type AccessSchemaJson struct {
	Roles               map[string]bool `json:"roles,omitempty"`
	Units               map[string]bool `json:"units,omitempty"`
	Users               map[string]bool `json:"users,omitempty"`
	Services            map[string]bool `json:"services,omitempty"`
	Stations            map[string]bool `json:"stations,omitempty"`
	Locations           map[string]bool `json:"locations,omitempty"`
	AssistentialCenters map[string]bool `json:"assistential_centers,omitempty"`
}

// RoleToDto converts a Role struct to a RoleJson struct.
func ModelToRoleJson(role *models.Role) *RoleJson {
	accessSchemaDto := AccessSchemaJson{
		Roles:               map[string]bool{},
		Units:               map[string]bool{},
		Users:               map[string]bool{},
		Services:            map[string]bool{},
		Stations:            map[string]bool{},
		Locations:           map[string]bool{},
		AssistentialCenters: map[string]bool{},
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
		Id:           utils.ParseInt64Sring(role.ID),
		RoleName:     role.RoleName,
		AccessSchema: accessSchemaDto,
		StRole:       role.StRole,
		Created_at:   role.CreatedAt.Format(time.RFC3339),
		Updated_at:   role.UpdatedAt.Format(time.RFC3339),
	}
}

func addPermissionsToDto(dtoPermissions *map[string]bool, permissions models.Permissions) {
	if len(permissions) > 0 {
		*dtoPermissions = make(map[string]bool)
		for k, v := range permissions {
			(*dtoPermissions)[k] = v
		}
	}
}

// addDtoPermissionsToRole add dto permissions to the role if the dto permissions are not empty
func addDtoPermissionsToRole(permissions *models.Permissions, dtoPermissions map[string]bool) error {
	if len(dtoPermissions) > 0 {
		*permissions = make(models.Permissions)
		for k, v := range dtoPermissions {

			(*permissions)[k] = v
		}
	}
	return nil
}

// DtoToRole converts a RoleJson struct to a Role struct.
func (dto *RoleJson) ToModel() models.Role {

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

	id := utils.ParseInt64(dto.Id)

	// stRole, err := strconv.Atoi(dto.StRole)
	// if err != nil {
	// 	return Role{}, fmt.Errorf("error converting st_role to integer: %w", err)
	// }

	createdAt, _ := time.Parse(time.RFC3339, dto.Created_at)
	// if err != nil {
	// 	return &models.Role{}, fmt.Errorf("error parsing created_at: %w", err)
	// }

	updatedAt, _ := time.Parse(time.RFC3339, dto.Updated_at)
	// if err != nil {
	// 	return &models.Role{}, fmt.Errorf("error parsing updated_at: %w", err)
	// }

	return models.Role{
		ID:           id,
		RoleName:     dto.RoleName,
		StRole:       dto.StRole,
		AccessSchema: accessSchema,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
