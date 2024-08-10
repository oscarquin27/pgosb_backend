package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type StationService interface {
	abstract_handler.AbstractCRUDService[models.Station]
}
