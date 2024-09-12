package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type RelevantServicesJson struct {
	Id                    string                        `json:"id"`
	RegionAreaDescription string                        `json:"region_area"`
	MissionCode           string                        `json:"mission_code"`
	AntaresId             string                        `json:"antares_id"`
	AntaresType           string                        `json:"antares_type"`
	AntaresDescription    string                        `json:"antares_description"`
	ServiceId             string                        `json:"service_id"`
	OprativeAreaName      []string                        `json:"operative_area_name"`
	ServiceDescription    string                        `json:"service_description"`
	ServiceDate           string                        `json:"service_date"`
	Units                 []string                      `json:"units"`
	Firefighters          []models.RelevantFirefighters        `json:"firefighters"`
	People                []models.RelevantPeople              `json:"people"`
	Infrastructures       []models.RelevantInfrastructure       `json:"infrastructures"`
	Vehicles              []models.RelevantVehicle             `json:"vehicles"`
	ServiceLocations      []models.RelevantServiceLocation     `json:"service_locations"`
	ServiceStations       []models.RelevantServiceStation      `json:"service_stations"`
	Centers               []models.RelevantCenter              `json:"centers"`
}



func ModelToRelevantServicesJson(r models.RelevantServices) *RelevantServicesJson {
	relevantService := &RelevantServicesJson{}

	relevantService.Id = utils.ParseInt64Sring(r.Id) // Assuming utils.ConvertFromInt4 converts int32 to string
	relevantService.MissionCode = *r.MissionCode
	relevantService.RegionAreaDescription = *r.RegionAreaDescription
	relevantService.AntaresId = utils.ConvertIntToString(*r.AntaresId)
	relevantService.AntaresType = *r.AntaresType
	relevantService.AntaresDescription = *r.AntaresDescription
	relevantService.ServiceId = utils.ConvertIntToString(*r.ServiceId)

	units := make([]string, len(r.Units))
	operativeAreaName := make([]string, len(r.OprativeAreaName))

	relevantService.OprativeAreaName = operativeAreaName
	relevantService.ServiceDescription = *r.ServiceDescription
	relevantService.ServiceDate = *r.ServiceDate
	relevantService.Units = units
	relevantService.Firefighters = r.Firefighters
	relevantService.People = r.People
	relevantService.Infrastructures = r.Infrastructures
	relevantService.Vehicles = r.Vehicles
	relevantService.ServiceLocations = r.ServiceLocations
	relevantService.ServiceStations = r.ServiceStations
	relevantService.Centers = r.Centers

	return relevantService
}

func (r *RelevantServicesJson) ToModel() models.RelevantServices {
	relevantService := models.RelevantServices{}

	antaresId := utils.ParseInt(r.AntaresId)
	serviceId := utils.ParseInt(r.ServiceId)


	relevantService.Id = utils.ParseInt64(r.Id) // Assuming utils.ConvertToPgTypeInt4 converts string to int32
	relevantService.RegionAreaDescription = &r.RegionAreaDescription
	relevantService.MissionCode = &r.MissionCode
	relevantService.AntaresId = &antaresId

	operativeAreaName := make([]*string, len(r.OprativeAreaName))
	units := make([]*string, len(r.Units))

	relevantService.OprativeAreaName = operativeAreaName
	relevantService.ServiceId = &serviceId
	relevantService.ServiceDescription = &r.ServiceDescription
	relevantService.ServiceDate = &r.ServiceDate
	relevantService.Units = units
	relevantService.Firefighters = r.Firefighters
	relevantService.People = r.People
	relevantService.Infrastructures = r.Infrastructures
	relevantService.Vehicles = r.Vehicles
	relevantService.ServiceLocations = r.ServiceLocations
	relevantService.ServiceStations = r.ServiceStations
	relevantService.Centers = r.Centers

	return relevantService
}