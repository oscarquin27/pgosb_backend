package services

import "fdms/src/models"

type MissionServiceService interface {
	GetAll() ([]models.MissionService, error)
	GetByMissionId(id int) ([]models.MissionService, error)
	Get(id int) (*models.MissionService, error)
	//GetAll() ([]models.MissionService, error)
	Create(user *models.MissionService) (*models.MissionService, error)
	Update(user *models.MissionService) error
	Delete(id int) error
}
