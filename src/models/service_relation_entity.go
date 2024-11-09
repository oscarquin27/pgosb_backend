package models

import (
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	ServiceRelationEntityTypeFirefighter = "firefighter"

	ServiceRelationEntityTypeInfrastructure = "infrastructure"

	ServiceRelationEntityTypePerson = "person"

	ServiceRelationEntityTypeUnit = "unit"

	ServiceRelationEntityTypeVehicle = "vehicle"
)

type ServiceRelationEntity struct {
	Id                      int64     `db:"id"`
	CreatedAt               time.Time `db:"created_at"`
	ServiceId               int64     `db:"service_id"`
	MissionId               int64     `db:"mission_id"`
	MissionEntityRelationId int64     `db:"mission_entity_relation_id"`
	Type                    string    `db:"-"`
}

func (s *ServiceRelationEntity) SetId(id int64) {
	s.Id = id
}

func (s *ServiceRelationEntity) GetNameArgs() pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":                         s.Id,
		"created_at":                 s.CreatedAt,
		"service_id":                 s.ServiceId,
		"mission_id":                 s.MissionId,
		"mission_entity_relation_id": s.MissionEntityRelationId,
	}
}
