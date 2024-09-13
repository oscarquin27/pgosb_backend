package api_models

import (
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionServiceJson struct {
	Id         string `json:"id"`
	MissionId  string `json:"mission_id"`
	AntaresId  string `json:"antares_id"`
	StationId  string `json:"station_id"`
	LocationId string `json:"location_id"`

	HealthCareCenterId string   `json:"center_id"`
	Units              []string `json:"units"`
	Bombers            []string `json:"bombers"`
	Summary            string   `json:"summary"`
	Description        string   `json:"description"`
	Unharmed           string   `json:"unharmed"`
	Injured            string   `json:"injured"`
	Transported        string   `json:"transported"`
	Deceased           string   `json:"deceased"`
	OperativeAreas     []string `json:"operative_areas"`

	ServiceDate       string `json:"service_date"`
	ManualServiceDate string `json:"manual_service_date"`
	IsImportant       bool   `json:"is_important"`

	SendingUserId   string `json:"sending_user_id"`
	ReceivingUserId string `json:"receiving_user_id"`
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
	service.LocationId = utils.ParseInt64SringPointer(s.LocationId)
	service.OperativeAreas = utils.ConvertFromTextArray(s.OperativeAreas)
	service.HealthCareCenterId = utils.ParseInt64SringPointer(s.HealthCareCenterId)

	if s.ManualServiceDate.Valid {
		service.ManualServiceDate = s.ManualServiceDate.Time.Format("02-01-2006 15:04:05")
	}

	if s.ServiceDate.Valid {
		service.ServiceDate = s.ServiceDate.Time.Format("02-01-2006 15:04:05")
	}
	service.IsImportant = s.IsImportant

	service.SendingUserId = utils.ParseInt64SringPointer(s.SendingUserId)
	service.ReceivingUserId = utils.ParseInt64SringPointer(s.ReceivingUserId)

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
	service.OperativeAreas = utils.ConvertToTextArray(s.OperativeAreas)

	unharmed := utils.ParseInt64(s.Unharmed)
	injured := utils.ParseInt64(s.Injured)
	transported := utils.ParseInt64(s.Transported)
	deceased := utils.ParseInt64(s.Deceased)
	stationId := utils.ParseInt64(s.StationId)
	centerId := utils.ParseInt64(s.HealthCareCenterId)

	locationId := utils.ParseInt64(s.LocationId)

	service.Unharmed = &unharmed
	service.Injured = &injured
	service.Transported = &transported
	service.Deceased = &deceased
	service.StationId = &stationId
	service.HealthCareCenterId = &centerId
	service.LocationId = &locationId
	manualServiceDate, err := time.Parse("02-01-2006 15:04:05", s.ManualServiceDate)

	if err == nil {
		service.ManualServiceDate.Time = manualServiceDate
		service.ManualServiceDate.Valid = true
	} else {
		logger.Warn().Err(err).Msg("Problema parseando manual service date")
	}

	// serviceDate, err := time.Parse("02-01-2006 15:04:05", s.ServiceDate)

	// if err == nil {
	// 	service.ServiceDate = serviceDate
	// } else {
	// 	logger.Warn().Err(err).Msg("Problema parseando service date")
	// }

	sendingUserId := utils.ParseInt64(s.SendingUserId)
	receivingUserId := utils.ParseInt64(s.ReceivingUserId)

	service.SendingUserId = &sendingUserId
	service.ReceivingUserId = &receivingUserId

	service.IsImportant = s.IsImportant

	return service
}
