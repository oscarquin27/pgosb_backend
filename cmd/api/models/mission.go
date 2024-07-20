package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionJson struct {
	Id        string `json:"id" binding:"required"`
	CreatedAt string `json:"created_at"`
	Code      string `json:"code"`
}

func ModelToMissionJson(s models.Mission) *MissionJson {
	mission := MissionJson{}

	mission.Id = utils.ConvertFromInt4(s.Id)
	mission.CreatedAt = utils.ConvertFromDate(s.CreatedAt)
	mission.Code = utils.ConvertFromText(s.Code)

	return &mission
}

func (s *MissionJson) ToModel() models.Mission {
	mission := models.Mission{}

	mission.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	mission.CreatedAt = utils.ConvertToPgTypeDate(s.CreatedAt)
	mission.Code = utils.ConvertToPgTypeText(s.Code)

	return mission
}