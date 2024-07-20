package services

import "fdms/src/models"

type MissionInfrastructureService interface {
	Get(id int) ([]models.MissionInfrastructure, error)
	Create(infra *models.MissionInfrastructure) error
	Update(infra *models.MissionInfrastructure) error
	Delete(id int) error
}