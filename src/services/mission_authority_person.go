package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionAuthorityPersonService interface {
	abstract_handler.AbstractCRUDService[models.MissionAuthorityPerson]
	GetByAuthorityId(id int64) *results.ResultWithValue[[]models.MissionAuthorityPerson]
}
