package api_models

import (
	"database/sql"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionServiceJson struct {
	Id        string `json:"id"`
	MissionId string `json:"mission_id"`
	AntaresId string `json:"antares_id"`

	StationId         string `json:"station_id"`
	LocationId        string `json:"location_id"`
	LocationDestinyId string `json:"location_destiny_id"`

	HealthCareCenterId string `json:"center_id"`

	Summary     string `json:"summary"`
	Description string `json:"description"`

	//Unharmed           string   `json:"unharmed"`

	ServiceDate       string `json:"service_date"`
	ManualServiceDate string `json:"manual_service_date"`
	IsImportant       bool   `json:"is_important"`

	SendingUserId   string `json:"sending_user_id"`
	ReceivingUserId string `json:"receiving_user_id"`

	Level         string `json:"level"`
	PeaceQuadrant string `json:"peace_quadrant"`

	CanceledReason string `json:"cancel_reason"`
	PendingForData bool   `json:"pending_for_data"`
}

func ModelToMissionServiceJson(s models.MissionService) *MissionServiceJson {
	service := MissionServiceJson{}

	service.Id = utils.ConvertFromInt4(s.Id)
	service.MissionId = utils.ConvertFromInt2(s.MissionId)
	service.AntaresId = utils.ConvertFromInt2(s.AntaresId)

	service.Summary = utils.ConvertFromText(s.Summary)
	service.Description = utils.ConvertFromText(s.Description)

	service.StationId = utils.ParseInt64StringPointer(s.StationId)
	service.LocationId = utils.ParseInt64StringPointer(s.LocationId)

	service.HealthCareCenterId = utils.ParseInt64StringPointer(s.HealthCareCenterId)

	if s.PeaceQuadrant.Valid {
		service.PeaceQuadrant = s.PeaceQuadrant.String
	}

	if s.ManualServiceDate.Valid {

		service.ManualServiceDate = s.ManualServiceDate.Time.Format("02-01-2006 15:04:05")
	}

	if s.ServiceDate.Valid {
		service.ServiceDate = s.ServiceDate.Time.Format("02-01-2006 15:04:05")
	}
	service.IsImportant = s.IsImportant

	service.SendingUserId = utils.ParseInt64StringPointer(s.SendingUserId)
	service.ReceivingUserId = utils.ParseInt64StringPointer(s.ReceivingUserId)

	if s.Level.Valid {
		service.Level = s.Level.String
	}

	service.LocationDestinyId = utils.ParseInt64StringPointer(s.LocationDestinyId)

	if s.CanceledReason.Valid {
		service.CanceledReason = s.CanceledReason.String
	}

	if s.PendingForData.Valid {
		service.PendingForData = s.PendingForData.Bool
	}

	return &service
}

func (s *MissionServiceJson) ToModel() models.MissionService {
	service := models.MissionService{}

	service.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	service.MissionId = utils.ConvertToPgTypeInt2(utils.ParseInt(s.MissionId))
	service.AntaresId = utils.ConvertToPgTypeInt2(utils.ParseInt(s.AntaresId))

	service.Summary = utils.ConvertToPgTypeText(s.Summary)
	service.Description = utils.ConvertToPgTypeText(s.Description)

	stationId := utils.ParseInt64(s.StationId)
	centerId := utils.ParseInt64(s.HealthCareCenterId)

	locationId := utils.ParseInt64(s.LocationId)

	service.StationId = &stationId
	service.HealthCareCenterId = &centerId
	service.LocationId = &locationId
	manualServiceDate, err := time.Parse("02-01-2006 15:04:05", s.ManualServiceDate)

	if s.PeaceQuadrant != "" {
		service.PeaceQuadrant = sql.NullString{String: s.PeaceQuadrant, Valid: true}
	}

	if err == nil {
		service.ManualServiceDate.Time = manualServiceDate
		service.ManualServiceDate.Valid = true
	} else {
		logger.Warn().Err(err).Msg("Problema parseando manual service date")
	}

	locationDestinyId := utils.ParseInt64(s.LocationDestinyId)
	service.LocationDestinyId = &locationDestinyId

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

	if s.Level != "" {
		service.Level = sql.NullString{String: s.Level, Valid: true}
	}

	service.PendingForData = sql.NullBool{Bool: s.PendingForData, Valid: true}

	return service
}
