package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionJson struct {
	Id             string   `json:"id"`
	CreatedAt      string   `json:"created_at"`
	Code           string   `json:"code"`
	Alias          string   `json:"alias"`
	OperativeAreas []string `json:"operative_areas,omitempty"`
	Summary        string   `json:"summary,omitempty"`
	Description    string   `json:"description,omitempty"`
	Unharmed       string   `json:"unharmed"`
	Injured        string   `json:"injured"`
	Transported    string   `json:"transported"`
	Deceased       string   `json:"deceased"`
	StationId      string   `json:"station_id"`
	LocationId     string   `json:"location_id"`
	//ServiceDate        string   `json:"service_date"`
	ManualMissionDate  string `json:"manual_mission_date"`
	IsImportant        bool   `json:"is_important"`
	HealthCareCenterId string `json:"center_id"`
	SendingUserId      string `json:"sending_user_id"`
	ReceivingUserId    string `json:"receiving_user_id"`
	LocationDestinyId  string `json:"location_destiny_id"`
	Level              string `json:"level"`
	PeaceQuadrant      string `json:"peace_quadrant"`
	CanceledReason     string `json:"cancel_reason"`
	PendingForData     bool   `json:"pending_for_data"`
}

func ModelToMissionJson(s models.Mission) *MissionJson {
	mission := MissionJson{
		Id: utils.ParseInt64String(s.Id),
	}

	if s.CreatedAt.Valid {
		mission.CreatedAt = s.CreatedAt.Time.Format("02-01-2006 15:04:05")
	}

	if s.Code.Valid {
		mission.Code = s.Code.String
	}

	if s.Alias.Valid {
		mission.Alias = s.Alias.String
	}

	op := make([]string, 0)
	for _, i := range s.OperativeAreas {
		if i.Valid {
			op = append(op, i.String)
		}
	}
	mission.OperativeAreas = op

	if s.Summary.Valid {
		mission.Summary = s.Summary.String
	}

	if s.Description.Valid {
		mission.Description = s.Description.String
	}

	if s.Unharmed.Valid {
		mission.Unharmed = utils.ParseInt64String(s.Unharmed.Int64)
	}

	if s.Injured.Valid {
		mission.Injured = utils.ParseInt64String(s.Injured.Int64)
	}

	if s.Transported.Valid {
		mission.Transported = utils.ParseInt64String(s.Transported.Int64)
	}

	if s.Deceased.Valid {
		mission.Deceased = utils.ParseInt64String(s.Deceased.Int64)
	}

	if s.StationId.Valid {
		mission.StationId = utils.ParseInt64String(s.StationId.Int64)
	}

	if s.LocationId.Valid {
		mission.LocationId = utils.ParseInt64String(s.LocationId.Int64)
	}

	//if s.ServiceDate.Valid {
	//	mission.ServiceDate = s.ServiceDate.Time.Format("02-01-2006 15:04:05")
	//}

	if s.ManualMissionDate.Valid {
		mission.ManualMissionDate = s.ManualMissionDate.Time.Format("02-01-2006 15:04:05")
	}

	mission.IsImportant = s.IsImportant.Bool

	if s.HealthCareCenterId.Valid {
		mission.HealthCareCenterId = utils.ParseInt64String(s.HealthCareCenterId.Int64)
	}

	if s.SendingUserId.Valid {
		mission.SendingUserId = utils.ParseInt64String(s.SendingUserId.Int64)
	}

	if s.ReceivingUserId.Valid {
		mission.ReceivingUserId = utils.ParseInt64String(s.ReceivingUserId.Int64)
	}

	if s.Level.Valid {
		mission.Level = s.Level.String
	}

	if s.PeaceQuadrant.Valid {
		mission.PeaceQuadrant = s.PeaceQuadrant.String
	}

	if s.LocationDestinyId.Valid {
		mission.LocationDestinyId = utils.ParseInt64String(s.LocationDestinyId.Int64)
	}

	if s.PendingForData.Valid {
		mission.PendingForData = s.PendingForData.Bool
	}

	if s.CanceledReason.Valid {
		mission.CanceledReason = s.CanceledReason.String
	}

	return &mission
}

func (s *MissionJson) ToModel() models.Mission {
	mission := models.Mission{
		Id: utils.ParseInt64(s.Id),
	}

	if s.CreatedAt != "" {
		if parsedTime, err := time.Parse("02-01-2006 15:04:05", s.CreatedAt); err == nil {
			mission.CreatedAt.Time = parsedTime
			mission.CreatedAt.Valid = true
		}
	}

	if s.Code != "" {
		mission.Code.String = s.Code
		mission.Code.Valid = true
	}

	if s.Alias != "" {
		mission.Alias.String = s.Alias
		mission.Alias.Valid = true
	}

	op := make([]sql.NullString, 0)
	for _, i := range s.OperativeAreas {
		val := sql.NullString{
			String: i,
			Valid:  true,
		}
		op = append(op, val)
	}
	mission.OperativeAreas = op

	if s.Summary != "" {
		mission.Summary.String = s.Summary
		mission.Summary.Valid = true
	}

	if s.Description != "" {
		mission.Description.String = s.Description
		mission.Description.Valid = true
	}

	if s.Unharmed != "" {
		mission.Unharmed.Int64 = utils.ParseInt64(s.Unharmed)
		mission.Unharmed.Valid = true
	}

	if s.Injured != "" {
		mission.Injured.Int64 = utils.ParseInt64(s.Injured)
		mission.Injured.Valid = true
	}

	if s.Transported != "" {
		mission.Transported.Int64 = utils.ParseInt64(s.Transported)
		mission.Transported.Valid = true
	}

	if s.Deceased != "" {
		mission.Deceased.Int64 = utils.ParseInt64(s.Deceased)
		mission.Deceased.Valid = true
	}

	if s.StationId != "" {
		mission.StationId.Int64 = utils.ParseInt64(s.StationId)
		mission.StationId.Valid = true
	}

	if s.LocationId != "" {
		mission.LocationId.Int64 = utils.ParseInt64(s.LocationId)
		mission.LocationId.Valid = true
	}

	//if s.ServiceDate != "" {
	//	if parsedTime, err := time.Parse("02-01-2006 15:04:05", s.ServiceDate); err == nil {
	//		mission.ServiceDate.Time = parsedTime
	//		mission.ServiceDate.Valid = true
	//	}
	//}

	if s.ManualMissionDate != "" {
		if parsedTime, err := time.Parse("02-01-2006 15:04:05", s.ManualMissionDate); err == nil {
			mission.ManualMissionDate.Time = parsedTime
			mission.ManualMissionDate.Valid = true
		}
	}

	mission.IsImportant.Bool = s.IsImportant
	mission.IsImportant.Valid = true

	if s.HealthCareCenterId != "" {
		mission.HealthCareCenterId.Int64 = utils.ParseInt64(s.HealthCareCenterId)
		mission.HealthCareCenterId.Valid = true
	}

	if s.SendingUserId != "" {
		mission.SendingUserId.Int64 = utils.ParseInt64(s.SendingUserId)
		mission.SendingUserId.Valid = true
	}

	if s.ReceivingUserId != "" {
		mission.ReceivingUserId.Int64 = utils.ParseInt64(s.ReceivingUserId)
		mission.ReceivingUserId.Valid = true
	}

	if s.Level != "" {
		mission.Level.String = s.Level
		mission.Level.Valid = true
	}

	if s.PeaceQuadrant != "" {
		mission.PeaceQuadrant.String = s.PeaceQuadrant
		mission.PeaceQuadrant.Valid = true
	}

	if s.LocationDestinyId != "" {
		mission.LocationDestinyId.Int64 = utils.ParseInt64(s.LocationDestinyId)
		mission.LocationDestinyId.Valid = true
	}

	if s.CanceledReason != "" {
		mission.CanceledReason.String = s.CanceledReason
		mission.CanceledReason.Valid = true
	}

	mission.PendingForData.Bool = s.PendingForData
	mission.PendingForData.Valid = true

	return mission
}
