package services

import "fdms/src/models"

type MIssionVehicleService interface {
	Get(id int) (*models.MissionVehicle, error)
	GetAll() ([]models.MissionVehicle, error)
	Create(user *models.MissionVehicle) error
	Update(user *models.MissionVehicle) error
	Delete(id int) error
}