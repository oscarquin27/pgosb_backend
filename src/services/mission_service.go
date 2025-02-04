package services

import (
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionServiceService interface {
	GetAll(isTemplate bool) ([]models.MissionService, error)

	GetAllMissionServiceSummary(isTemplate bool) ([]models.MissionServiceSummary, error)

	GetRelevantServices(id string, isTemplate bool) ([]models.RelevantServices, error)

	GetRelevantMissions(id string, isTemplate bool) ([]models.RelevantServices, error)

	GetByMissionId(id int64, isTemplate bool) ([]models.MissionService, error)

	Get(id int64, isTemplate bool) (*models.MissionService, error)

	Create(user *models.MissionService, isTemplate bool) (*models.MissionService, error)

	Update(user *models.MissionService, isTemplate bool) error

	Delete(id int64, isTemplate bool) error

	GetUnits(id int64, isTemplate bool) *results.ResultWithValue[[]models.UnitSimple]

	GetUsers(id int64, isTemplate bool) *results.ResultWithValue[[]models.MissionUserService]
}
