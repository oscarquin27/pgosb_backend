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

type MissionFirefighterUserJson struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	User_name    string `json:"user_name"`
	Rank         string `json:"rank"`
	PersonalCode string `json:"personal_code"`
	Legal_id     string `json:"legal_id"`
	MissionId    string `json:"mission_id"`
}

func ModelToMissionFirefighterUserJson(s *models.MissionFirefighterUser) *MissionFirefighterUserJson {
	service := MissionFirefighterUserJson{}
	service.Id = utils.ParseInt64String(s.Id)
	service.UserId = utils.ParseInt64String(s.UserId)
	service.Name = s.Name

	service.User_name = utils.GetStringFromPointer(s.User_name)
	service.Rank = s.Rank
	service.PersonalCode = s.PersonalCode
	service.Legal_id = s.Legal_id
	service.MissionId = s.MissionId

	return &service
}

func (s *MissionFirefighterUserJson) ToModel() models.MissionFirefighterUser {
	service := models.MissionFirefighterUser{}
	service.Id = utils.ParseInt64(s.Id)
	service.UserId = utils.ParseInt64(s.UserId)
	service.Name = s.Name
	service.User_name = &s.User_name
	service.Rank = s.Rank
	service.PersonalCode = s.PersonalCode
	service.Legal_id = s.Legal_id
	service.MissionId = s.MissionId

	return service
}
