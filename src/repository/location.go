package repository

import "fdms/src/models"

type LocationsService interface {
	GetState(id int64) (*models.State, error)
	GetAllStates() ([]models.State, error)
	CreateState(models *models.State) error
	UpdateState(models *models.State) error
	DeleteState(id int64) error
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
	GetStation(id int64) (*models.Station, error)
	GetAllStations() ([]models.Station, error)
	CreateStation(models *models.Station) error
	UpdateStation(models *models.Station) error
	DeleteStation(id int64) error
}
