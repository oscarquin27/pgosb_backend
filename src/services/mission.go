package services

import "fdms/src/models"

type MissionService interface {
	Get(id int) (*models.Mission, error)
	GetAll() ([]models.Mission, error)
	GetAllMissionSummary() ([]models.MissionSummary, error)

	Create(user *models.Mission) (*models.Mission, error)
	Update(user *models.Mission) error
	Delete(id int) error
}
