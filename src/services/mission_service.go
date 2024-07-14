package services

import "fdms/src/models"

type ServiceService interface {
	Get(id int) ([]models.MissionService, error)
	GetAll() ([]models.MissionService, error)
	Create(user *models.MissionService) error
	Update(user *models.MissionService) error
	Delete(id int) error
}
