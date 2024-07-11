package services

import "fdms/src/models"

type VehicleService interface {
	Get(id int64) (*models.Vehicle, error)
	GetAll() ([]models.Vehicle, error)
	Create(user *models.Vehicle) error
	Update(user *models.Vehicle) error
	Delete(id int64) error
}
