package services

import "fdms/src/models"

type MissionService interface {
	Get(id int64, isTemplrate bool) (*models.Mission, error)
	GetAll(isTemplate bool) ([]models.Mission, error)
	GetAllMissionSummary(isTemplate bool) ([]models.MissionSummary, error)

	Create(user *models.Mission, isTemplate bool) (*models.Mission, error)
	Update(user *models.Mission, isTemplate bool) error
	Delete(id int64, isTemplate bool) error
	GetSpecialOperationList() ([]string, error)
}
