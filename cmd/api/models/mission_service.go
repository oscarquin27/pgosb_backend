package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionServiceJson struct {
	Id                 string   `json:"id"`
	MissionId          string   `json:"mission_id"`
	AntaresId          string   `json:"antares_id"`
	StationId          string   `json:"station_id"`
	HealthCareCenterId string   `json:"center_id"`
	Units              []string `json:"units"`
	Bombers            []string `json:"bombers"`
	Summary            string   `json:"summary"`
	Description        string   `json:"description"`
	Unharmed           string   `json:"unharmed"`
	Injured            string   `json:"injured"`
	Transported        string   `json:"transported"`
	Deceased           string   `json:"deceased"`
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

	service.Unharmed = utils.ParseInt64SringPointer(s.Unharmed)
	service.Injured = utils.ParseInt64SringPointer(s.Injured)
	service.Transported = utils.ParseInt64SringPointer(s.Transported)
	service.Deceased = utils.ParseInt64SringPointer(s.Deceased)
	service.StationId = utils.ParseInt64SringPointer(s.StationId)
	service.HealthCareCenterId = utils.ParseInt64SringPointer(s.HealthCareCenterId)

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

	unharmed := utils.ParseInt64(s.Unharmed)
	injured := utils.ParseInt64(s.Injured)
	transported := utils.ParseInt64(s.Transported)
	deceased := utils.ParseInt64(s.Deceased)
	stationId := utils.ParseInt64(s.StationId)
	centerId := utils.ParseInt64(s.HealthCareCenterId)

	service.Unharmed = &unharmed
	service.Injured = &injured
	service.Transported = &transported
	service.Deceased = &deceased
	service.StationId = &stationId
	service.HealthCareCenterId = &centerId

	return service
}
