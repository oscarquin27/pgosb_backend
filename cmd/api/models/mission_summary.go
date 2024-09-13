package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionSummaryJson struct {
	Id              string `json:"id"`
	Alias           string `json:"alias"`
	CreatedAt       string `json:"created_at"`
	NumServices     string `json:"num_services"`
	NumFireFighters string `json:"num_firefighters"`
	NumVehicles     string `json:"num_vehicles"`
	Unharmed        string `json:"unharmed"`
	Injured         string `json:"injured"`
	Transported     string `json:"transported"`
	Deceased        string `json:"deceased"`
	Code            string `json:"code"`
	OperativeAreas  []string `json:"operative_areas"`
}

func ModelToMissionSummaryJson(s models.MissionSummary) *MissionSummaryJson {
	service := MissionSummaryJson{}

	service.Id = utils.ParseInt64Sring(s.Id)

	if s.Alias.Valid {
		service.Alias = s.Alias.String
	}

	if s.CreatedAt.Valid {
		service.CreatedAt = s.CreatedAt.Time.Format("02-01-2006 15:04:05")
	}

	if s.NumServices.Valid {
		service.NumServices = utils.ParseInt64Sring(s.NumServices.Int64)
	}

	if s.NumVehicles.Valid {
		service.NumVehicles = utils.ParseInt64Sring(s.NumVehicles.Int64)
	}

	if s.NumFireFighters.Valid {
		service.NumFireFighters = utils.ParseInt64Sring(s.NumFireFighters.Int64)
	}

	if s.Unharmed.Valid {
		service.Unharmed = utils.ParseInt64Sring(s.Unharmed.Int64)
	}

	if s.Injured.Valid {
		service.Injured = utils.ParseInt64Sring(s.Injured.Int64)
	}

	if s.Transported.Valid {
		service.Transported = utils.ParseInt64Sring(s.Transported.Int64)
	}

	if s.Deceased.Valid {
		service.Deceased = utils.ParseInt64Sring(s.Deceased.Int64)
	}

	if s.Code.Valid {
		service.Code = s.Code.String
	}

	op := make([]string, 0)
	
	for _, i := range s.OperativeAreas {
		i.Valid = true
		op = append(op, i.String)
	}

	service.OperativeAreas = op

	return &service
}

func (s *MissionSummaryJson) ToModel() models.MissionSummary {
	service := models.MissionSummary{}

	service.Id = utils.ParseInt64(s.Id)

	if s.Alias != "" {
		service.Alias.String = s.Alias
		service.Alias.Valid = true
	}

	if s.CreatedAt != "" {
		service.CreatedAt.Time, _ = time.Parse("02-01-2006 15:04:05", s.CreatedAt)
		service.CreatedAt.Valid = true
	}

	service.Alias.String = s.Alias
	service.Alias.Valid = true

	if s.NumServices != "" {
		service.NumServices.Int64 = utils.ParseInt64(s.NumServices)
		service.NumServices.Valid = true
	}

	if s.NumVehicles != "" {
		service.NumVehicles.Int64 = utils.ParseInt64(s.NumVehicles)
		service.NumVehicles.Valid = true
	}

	if s.NumFireFighters != "" {
		service.NumFireFighters.Int64 = utils.ParseInt64(s.NumFireFighters)
		service.NumFireFighters.Valid = true
	}

	if s.Unharmed != "" {
		service.Unharmed.Int64 = utils.ParseInt64(s.Unharmed)
		service.Unharmed.Valid = true
	}

	if s.Injured != "" {
		service.Injured.Int64 = utils.ParseInt64(s.Injured)
		service.Injured.Valid = true
	}

	if s.Transported != "" {
		service.Transported.Int64 = utils.ParseInt64(s.Transported)
		service.Transported.Valid = true
	}

	if s.Deceased != "" {
		service.Deceased.Int64 = utils.ParseInt64(s.Deceased)
		service.Deceased.Valid = true
	}

	if s.Code != "" {
		service.Code.String = s.Code
		service.Code.Valid = true
	}

	op := make([]sql.NullString, 0)
	
	for _, i := range s.OperativeAreas {
		val := sql.NullString{
			String: i,
			Valid: true,
		}
		op = append(op, val)
	}

	service.OperativeAreas = op

	return service
}
