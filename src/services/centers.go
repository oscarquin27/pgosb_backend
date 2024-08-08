package services

import "fdms/src/models"

type CentersService interface {
	Get(id int) (*models.Center, error)
	GetAll() ([]models.Center, error)
	Create(infra *models.Center) error
	Update(infra *models.Center) error
	Delete(id int) error
}
