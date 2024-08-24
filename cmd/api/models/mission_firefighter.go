package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionFirefighterJson struct {
	Id            string `json:"id"`
	MissionId     string `json:"mission_id"`
	ServiceId     string `json:"service_id"`
	FireFighterId string `json:"firefigther_id"`
	ServiceRol    string `json:"service_role"`
}

func ModelToMissionFirefighterJson(s models.MissionFirefighter) *MissionFirefighterJson {
	service := MissionFirefighterJson{}
	service.Id = utils.ParseInt64Sring(s.Id)
	service.MissionId = utils.ParseInt64Sring(s.MissionId)
	service.ServiceId = utils.ParseInt64Sring(s.ServiceId)
	service.FireFighterId = utils.ParseInt64Sring(s.FireFighterId)
	service.ServiceRol = utils.GetStringFromPointer(s.ServiceRol)

	return &service
}

func (s *MissionFirefighterJson) ToModel() models.MissionFirefighter {
	service := models.MissionFirefighter{}

	service.Id = utils.ParseInt64(s.Id)
	service.MissionId = utils.ParseInt64(s.MissionId)
	service.ServiceId = utils.ParseInt64(s.ServiceId)
	service.FireFighterId = utils.ParseInt64(s.FireFighterId)
	service.ServiceRol = &s.ServiceRol

	return service
}
