package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type MissionAuthorityService interface {
	abstract_handler.AbstractCRUDService[models.MissionAuthority]
	GetByMissionId(id int64) *results.ResultWithValue[[]models.MissionAuthority]
	GetSummaryByMissionId(id int64) *results.ResultWithValue[[]models.MissionAuthoritySummary]
}
