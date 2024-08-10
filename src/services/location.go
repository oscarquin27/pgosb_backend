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

type LocationsService interface {
	GetCity(id int64) (*models.City, error)
	GetAllCity() ([]models.City, error)
	CreateCity(models *models.City) error
	UpdateCity(models *models.City) error
	DeleteCity(id int64) error

	GetMunicipality(id int64) (*models.Municipality, error)
	GetAllMunicipality() ([]models.Municipality, error)
	CreateMunicipality(models *models.Municipality) error
	UpdateMunicipality(models *models.Municipality) error
	DeleteMunicipality(id int64) error

	GetParish(id int64) (*models.Parish, error)
	GetAllParish() ([]models.Parish, error)
	CreateParish(models *models.Parish) error
	UpdateParish(models *models.Parish) error
	DeleteParish(id int64) error
}
