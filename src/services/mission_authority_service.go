package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionAuthorityRelateService interface {
	abstract_handler.AbstractCRUDService[models.MissionAuthorityService]
	GetByServiceId(id int64) *results.ResultWithValue[[]models.MissionAuthorityServiceSummary]
}
