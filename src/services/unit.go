package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type UnitService interface {
	abstract_handler.AbstractCRUDService[models.Unit]
	GetAllSimple() *results.ResultWithValue[[]models.UnitSimple]
}
