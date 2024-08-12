package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type StateService interface {
	abstract_handler.AbstractCRUDService[models.State]
}

type MunicipalityService interface {
	abstract_handler.AbstractCRUDService[models.Municipality]
}

type ParishService interface {
	abstract_handler.AbstractCRUDService[models.Parish]
}

type SectorService interface {
	abstract_handler.AbstractCRUDService[models.Sector]
}

type LocationsService interface {
	GetCity(id int64) (*models.City, error)
	GetAllCity() ([]models.City, error)
	CreateCity(models *models.City) error
	UpdateCity(models *models.City) error
	DeleteCity(id int64) error
}
