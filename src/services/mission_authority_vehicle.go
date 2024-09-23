package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionAuthorityVehicleService interface {
	abstract_handler.AbstractCRUDService[models.MissionAuthorityVehicle]
	GetByAuthorityId(id int64) *results.ResultWithValue[[]models.MissionAuthorityVehicle]
}
