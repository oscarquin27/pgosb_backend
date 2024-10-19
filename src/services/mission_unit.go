package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type MissionUnitService interface {
	abstract_handler.AbstractCRUDService[models.MissionUnit]
	GetByMissionId(id int) ([]models.MissionUnitSummary, error)
}
