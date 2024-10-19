package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionUnitJson struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	MissionId string `json:"mission_id"`
	UnitId    string `json:"unit_id"`
}

func ModelToMissionUnitJson(s *models.MissionUnit) *MissionUnitJson {
	service := MissionUnitJson{}
	service.Id = utils.ParseInt64String(s.Id)
	service.MissionId = utils.ParseInt64String(s.MissionId)

	service.UnitId = utils.ParseInt64String(s.UnitId)

	if s.CreatedAt.Valid {
		service.CreatedAt = s.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &service
}

func (s *MissionUnitJson) ToModel() models.MissionUnit {
	service := models.MissionUnit{}

	service.Id = utils.ParseInt64(s.Id)
	service.MissionId = utils.ParseInt64(s.MissionId)
	service.UnitId = utils.ParseInt64(s.UnitId)

	return service
}

type MissionUnitSummaryJson struct {
	Id        string `json:"id"`
	MissionId string `json:"mission_id"`
	UnitId    string `json:"unit_id"`
	Plate     string `json:"plate"`
	Alias     string `json:"alias"`
	UnitType  string `json:"unit_type"`
}

func ModelToMissionUnitSummaryJson(s *models.MissionUnitSummary) *MissionUnitSummaryJson {
	service := MissionUnitSummaryJson{}
	service.Id = utils.ParseInt64String(s.Id)
	service.MissionId = utils.ParseInt64String(s.MissionId)
	service.UnitId = utils.ParseInt64String(s.UnitId)

	if s.Plate.Valid {
		service.Plate = s.Plate.String
	}
	if s.Alias.Valid {
		service.Alias = s.Alias.String
	}
	if s.UnitType.Valid {
		service.UnitType = s.UnitType.String
	}

	return &service

}

func (s *MissionUnitSummaryJson) ToModel() models.MissionUnitSummary {
	service := models.MissionUnitSummary{}

	service.Id = utils.ParseInt64(s.Id)
	service.MissionId = utils.ParseInt64(s.MissionId)
	service.UnitId = utils.ParseInt64(s.UnitId)
	service.Plate = sql.NullString{String: s.Plate, Valid: s.Plate != ""}
	service.Alias = sql.NullString{String: s.Alias, Valid: s.Alias != ""}
	service.UnitType = sql.NullString{String: s.UnitType, Valid: s.UnitType != ""}

	return service
}
