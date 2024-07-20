package services

import "fdms/src/models"

type UnitService interface {
	Get(id int64) (*models.Unit, error)
	GetUnitTypes() ([]string, error)
	GetAll() ([]models.Unit, error)
	Create(user *models.Unit) error
	Update(user *models.Unit) error
	Delete(id int64) error
}
