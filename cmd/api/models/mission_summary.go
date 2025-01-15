package api_models

import (
	"database/sql"
	"fdms/src/models"
	"fdms/src/utils"
	"time"
)

type MissionSummaryJson struct {
	Id                   string   `json:"id"`
	Code                 string   `json:"code"`
	Alias                string   `json:"alias"`
	AntaresId            string   `json:"antares_id"`
	AntaresDescription   string   `json:"antares_description"`
	NumServices          string   `json:"num_services"`
	IsImportant          bool     `json:"is_important"`
	StationName          string   `json:"station_name"`
	Commander            string   `json:"commander"`
	NumFirefighters      string   `json:"num_firefighters"`
	FirefighterArray     []string `json:"firefighter_array"`
	NumVehicles          string   `json:"num_vehicles"`
	UnitArray            []string `json:"unit_array"`
	OperativesAreas      []string `json:"operative_areas"`
	Description          string   `json:"description"`
	Level                string   `json:"level"`
	PeaceQuadrant        string   `json:"peace_quadrant"`
	PendingForData       bool     `json:"pending_for_data"`
	CancelReason         string   `json:"cancel_reason"`
	CreatedAt            string   `json:"created_at"`
	ManualMissionDate    string   `json:"manual_mission_date"`
	Unharmed             string   `json:"unharmed"`
	Injured              string   `json:"injured"`
	Transported          string   `json:"transported"`
	Deceased             string   `json:"deceased"`
	LocationId           string   `json:"location_id"`
	LocationDestinyId    string   `json:"location_destiny_id"`
	State				 string	  `json:"state"`
	Municipality		 string	  `json:"municipality"`
	Parish				 string	  `json:"parish"`
	NumAuthorities       string   `json:"num_authorities"`
	NumAuthorityServices string   `json:"num_authority_services"`
	NumAuthorityPerson   string   `json:"num_authority_person"`
	NumAuthorityVehicle  string   `json:"num_authority_vehicle"`
	StationId            string   `json:"station_id"`
}

func ModelToMissionSummaryJson(s models.MissionSummary) *MissionSummaryJson {
	service := MissionSummaryJson{}

	service.Id = utils.ParseInt64String(s.Id)

	if s.Description.Valid {
		service.Description = s.Description.String
	}

	if s.Level.Valid {
		service.Level = s.Level.String
	}

	if s.PeaceQuadrant.Valid {
		service.PeaceQuadrant = s.PeaceQuadrant.String
	}

	service.PendingForData = s.PendingForData

	if s.CancelReason.Valid {
		service.CancelReason = s.CancelReason.String
	}

	if s.AntaresDescription.Valid {

		service.AntaresDescription = s.AntaresDescription.String
	}

	if s.Alias.Valid {

		service.Alias = s.Alias.String
	}

	if s.CreatedAt.Valid {
		service.CreatedAt = s.CreatedAt.Time.Format("02-01-2006 15:04:05")
	}

	if s.NumServices.Valid {
		service.NumServices = utils.ParseInt64String(s.NumServices.Int64)
	}

	if s.NumVehicles.Valid {
		service.NumVehicles = utils.ParseInt64String(s.NumVehicles.Int64)
	}

	if s.NumFirefighters.Valid {
		service.NumFirefighters = utils.ParseInt64String(s.NumFirefighters.Int64)
	}

	if s.Unharmed.Valid {
		service.Unharmed = utils.ParseInt64String(s.Unharmed.Int64)
	}

	if s.Injured.Valid {
		service.Injured = utils.ParseInt64String(s.Injured.Int64)
	}

	if s.Transported.Valid {
		service.Transported = utils.ParseInt64String(s.Transported.Int64)
	}

	if s.Deceased.Valid {
		service.Deceased = utils.ParseInt64String(s.Deceased.Int64)
	}

	if s.Code.Valid {
		service.Code = s.Code.String
	}

	if s.StationId.Valid {
		service.StationId = utils.ParseInt64String(s.StationId.Int64)
	}

	if s.State.Valid {
		service.State = s.State.String
	}

	if s.Municipality.Valid {
		service.Municipality = s.Municipality.String
	}

	if s.Parish.Valid {
		service.Parish = s.Parish.String
	}

	service.IsImportant = s.IsImportant

	op := make([]string, 0)

	for _, i := range s.OperativesAreas {
		if i.Valid {
			op = append(op, i.String)
		}
	}

	un := make([]string, 0)

	for _, i := range s.UnitArray {
		if i.Valid {
			un = append(un, i.String)
		}
	}

	fa := make([]string, 0)

	for _, i := range s.FirefighterArray {
		if i.Valid {
			fa = append(fa, i.String)
		}
	}

	if s.NumAuthorities.Valid {
		service.NumAuthorities = utils.ParseInt64String(s.NumAuthorities.Int64)
	}

	if s.NumAuthorityServices.Valid {
		service.NumAuthorityServices = utils.ParseInt64String(s.NumAuthorityServices.Int64)
	}

	if s.NumAuthorityPerson.Valid {
		service.NumAuthorityPerson = utils.ParseInt64String(s.NumAuthorityPerson.Int64)
	}

	if s.NumAuthorityVehicle.Valid {
		service.NumAuthorityVehicle = utils.ParseInt64String(s.NumAuthorityVehicle.Int64)
	}

	service.OperativesAreas = op
	service.UnitArray = un
	service.FirefighterArray = fa

	if s.ManualMissionDate.Valid {
		service.ManualMissionDate = s.ManualMissionDate.Time.Format("02-01-2006 15:04:05")

	}

	if s.StationName.Valid {
		service.StationName = s.StationName.String
	}

	if s.AntaresId.Valid {
		service.AntaresId = utils.ParseInt64String(s.AntaresId.Int64)
	}

	if s.AntaresDescription.Valid {
		service.AntaresDescription = s.AntaresDescription.String
	}

	if s.LocationId.Valid {
		service.LocationId = utils.ParseInt64String(s.LocationId.Int64)
	}

	if s.LocationDestinyId.Valid {
		service.LocationDestinyId = utils.ParseInt64String(s.LocationDestinyId.Int64)
	}

	if s.Commander.Valid {
		service.Commander = s.Commander.String
	}

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

	if s.NumFirefighters != "" {
		service.NumFirefighters.Int64 = utils.ParseInt64(s.NumFirefighters)
		service.NumFirefighters.Valid = true
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

	if s.NumAuthorities != "" {
		service.NumAuthorities.Int64 = utils.ParseInt64(s.NumAuthorities)
		service.NumAuthorities.Valid = true
	}

	if s.NumAuthorityServices != "" {
		service.NumAuthorityServices.Int64 = utils.ParseInt64(s.NumAuthorityServices)
		service.NumAuthorityServices.Valid = true
	}

	if s.NumAuthorityPerson != "" {
		service.NumAuthorityPerson.Int64 = utils.ParseInt64(s.NumAuthorityPerson)
		service.NumAuthorityPerson.Valid = true
	}

	if s.NumAuthorityPerson != "" {
		service.NumAuthorityPerson.Int64 = utils.ParseInt64(s.NumAuthorityPerson)
		service.NumAuthorityPerson.Valid = true
	}

	if s.StationId != "" {
		service.StationId.Int64 = utils.ParseInt64(s.StationId)
		service.StationId.Valid = true
	}

	op := make([]sql.NullString, 0)

	for _, i := range s.OperativesAreas {
		val := sql.NullString{
			String: i,
			Valid:  true,
		}
		op = append(op, val)
	}

	service.OperativesAreas = op

	return service
}
