package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionLocationService interface {
	abstract_handler.AbstractCRUDService[models.MissionLocation]
	GetLocationsByServiceId(id int64) *results.ResultWithValue[[]models.MissionLocation]
}
