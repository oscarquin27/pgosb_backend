package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type MissionFirefighterService interface {
	abstract_handler.AbstractCRUDService[models.MissionFirefighter]
}
