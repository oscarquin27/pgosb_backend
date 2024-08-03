package services

import "fdms/src/models"

type AntaresService interface {
	GetAll() ([]models.Antares, error)
}
