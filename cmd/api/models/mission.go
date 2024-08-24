package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
	"fdms/src/utils/date_utils"
)

type MissionJson struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Code      string `json:"code"`
	Alias     string `json:"alias"`
}

func ModelToMissionJson(s models.Mission) *MissionJson {
	mission := MissionJson{}

	mission.Id = utils.ConvertFromInt4(s.Id)

	createdDate := s.CreatedAt.Time.Format(date_utils.CompleteFormatDate)
	mission.CreatedAt = createdDate
	mission.Code = utils.ConvertFromText(s.Code)

	mission.Alias = utils.GetStringFromPointer(s.Alias)

	return &mission
}

func (s *MissionJson) ToModel() models.Mission {
	mission := models.Mission{}

	mission.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	mission.CreatedAt = utils.ConvertToPgTypeDate(s.CreatedAt)
	mission.Code = utils.ConvertToPgTypeText(s.Code)
	mission.Alias = &s.Alias
	return mission
}
