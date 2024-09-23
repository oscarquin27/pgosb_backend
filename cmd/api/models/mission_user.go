package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionUserServiceJson struct {
	Id           string `json:"id"`
	ServiceId    string `json:"service_id"`
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	Rank         string `json:"rank"`
	PersonalCode string `json:"personal_code"`
	LegalId      string `json:"legal_id"`
	ServiceRole  string `json:"service_role"`
}

func ModelToMissionUserServiceJson(s models.MissionUserService) *MissionUserServiceJson {
	service := MissionUserServiceJson{}

	service.Id = utils.ParseInt64String(s.Id)
	service.ServiceId = utils.ParseInt64String(s.ServiceId)

	if s.UserId.Valid {
		service.UserId = utils.ParseInt64String(s.UserId.Int64)
	}

	if s.Name.Valid {
		service.Name = s.Name.String
	}

	if s.UserName.Valid {
		service.UserName = s.UserName.String
	}

	if s.Rank.Valid {
		service.Rank = s.Rank.String
	}

	if s.PersonalCode.Valid {
		service.PersonalCode = s.PersonalCode.String
	}

	if s.LegalId.Valid {
		service.LegalId = s.LegalId.String
	}

	if s.ServiceRole.Valid {
		service.ServiceRole = s.ServiceRole.String
	}

	return &service
}

func (s *MissionUserServiceJson) ToModel() models.MissionService {
	service := models.MissionService{}

	return service
}
