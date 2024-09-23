package api_models

import (
	"fdms/src/models"
)

type MissionAuthorityServiceJson struct {
	Id          int64 `db:"id"`
	MissionId   int64 `db:"mission_id"`
	ServiceId   int64 `db:"service_id"`
	AuthorityId int64 `db:"authority_id"`
}

func (s *MissionAuthorityServiceJson) ToModel() models.MissionAuthorityService {
	return models.MissionAuthorityService{

		Id:          s.Id,
		MissionId:   s.MissionId,
		ServiceId:   s.ServiceId,
		AuthorityId: s.AuthorityId,
	}
}

func ModelToMissionAuthorityServiceJson(model *models.MissionAuthorityService) *MissionAuthorityServiceJson {
	return &MissionAuthorityServiceJson{
		Id:          model.Id,
		MissionId:   model.MissionId,
		ServiceId:   model.ServiceId,
		AuthorityId: model.AuthorityId,
	}
}
