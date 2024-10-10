package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionFirefighterJson struct {
	Id         string `json:"id"`
	MissionId  string `json:"mission_id"`
	ServiceId  string `json:"service_id"`
	UserId     string `json:"user_id"`
	ServiceRol string `json:"service_role"`
}

func ModelToMissionFirefighterJson(s *models.MissionFirefighter) *MissionFirefighterJson {
	service := MissionFirefighterJson{}
	service.Id = utils.ParseInt64String(s.Id)
	service.MissionId = utils.ParseInt64String(s.MissionId)

	service.UserId = utils.ParseInt64String(s.UserId)
	service.ServiceRol = utils.GetStringFromPointer(s.ServiceRol)

	return &service
}

func (s *MissionFirefighterJson) ToModel() models.MissionFirefighter {
	service := models.MissionFirefighter{}

	service.Id = utils.ParseInt64(s.Id)
	service.MissionId = utils.ParseInt64(s.MissionId)
	service.UserId = utils.ParseInt64(s.UserId)
	service.ServiceRol = &s.ServiceRol

	return service
}
