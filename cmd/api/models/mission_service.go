package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionServiceJson struct {
	Id          string   `json:"id" binding:"required"`
	MissionId   string   `json:"mission_id"`
	AntaresId   string   `json:"antares_id"`
	Units       []string `json:"units"`
	Bombers     []string `json:"bombers"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
}

func ModelToMissionServiceJson(s models.MissionService) *MissionServiceJson {
	service := MissionServiceJson{}

	service.Id = utils.ConvertFromInt4(s.Id)
	service.MissionId = utils.ConvertFromInt2(s.MissionId)
	service.AntaresId = utils.ConvertFromInt2(s.AntaresId)
	service.Units = utils.ConvertFromInt2Array(s.Units)
	service.Bombers = utils.ConvertFromInt2Array(s.Bombers)
	service.Summary = utils.ConvertFromText(s.Summary)
	service.Description = utils.ConvertFromText(s.Description)
	
	return &service
}


func (s *MissionServiceJson) ToModel() models.MissionService {
    service := models.MissionService{}

	service.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	service.MissionId = utils.ConvertToPgTypeInt2(utils.ParseInt(s.MissionId))
	service.AntaresId = utils.ConvertToPgTypeInt2(utils.ParseInt(s.AntaresId))
	service.Units = utils.ConvertToInt2Array(s.Units)
	service.Bombers = utils.ConvertToInt2Array(s.Bombers)
	service.Summary = utils.ConvertToPgTypeText(s.Summary)
	service.Description = utils.ConvertToPgTypeText(s.Description)

	return service
}