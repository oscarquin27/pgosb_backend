package services

import (
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionServiceService interface {
	GetAll() ([]models.MissionService, error)
	GetAllMissionServiceSummary() ([]models.MissionServiceSummary, error)
	GetRelevantServices(id string) ([]models.RelevantServices, error)
	GetRelevantMissions(id string) ([]models.RelevantServices, error)
	GetByMissionId(id int64) ([]models.MissionService, error)
	Get(id int64) (*models.MissionService, error)
	//GetAll() ([]models.MissionService, error)
	Create(user *models.MissionService) (*models.MissionService, error)
	Update(user *models.MissionService) error
	Delete(id int64) error
	GetUnits(id int64) *results.ResultWithValue[[]models.UnitSimple]
	GetUsers(id int64) *results.ResultWithValue[[]models.MissionUserService]
}
