package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type HealthcareCenterService interface {
	abstract_handler.AbstractCRUDService[models.HealthcareCenter]
}
