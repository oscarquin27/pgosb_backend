package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type RelevantServicesJson struct {
	Id                    string                           `json:"id"`
	RegionAreaDescription string                           `json:"region_area"`
	MissionCode           string                           `json:"mission_code"`
	Antares				  []models.RelevantAntares 		   `json:"antares"`
	ServiceId             string                           `json:"service_id"`
	OprativeAreaName      []string                         `json:"operative_area_name"`
	ServiceDescription    string                           `json:"service_description"`
	ServiceDate           string                           `json:"service_date"`
	Unharmed              string                           `json:"unharmed"`
	Injured               string                           `json:"injured"`
	Transported           string                           `json:"transported"`
	Deceased              string                           `json:"deceased"`
	Units                 []string                         `json:"units"`
	Firefighters          []models.RelevantFirefighters    `json:"firefighters"`
	People                []models.RelevantPeople          `json:"people"`
	Infrastructures       []models.RelevantInfrastructure  `json:"infrastructures"`
	Vehicles              []models.RelevantVehicle         `json:"vehicles"`
	ServiceLocations      []models.RelevantServiceLocation `json:"service_locations"`
	ServiceStations       []models.RelevantServiceStation  `json:"service_stations"`
	Centers               []models.RelevantCenter          `json:"centers"`
	IsImportant           bool                             `json:"is_important"`
}

func ModelToRelevantServicesJson(r models.RelevantServices) *RelevantServicesJson {
	relevantService := &RelevantServicesJson{}

	relevantService.Id = utils.ParseInt64String(r.Id) // Assuming utils.ConvertFromInt4 converts int32 to string
	relevantService.MissionCode = utils.GetStringFromPointer(r.MissionCode)
	relevantService.RegionAreaDescription = utils.GetStringFromPointer(r.RegionAreaDescription)

	if r.ServiceId != nil {
		relevantService.ServiceId = utils.ConvertInt64ToString(*r.ServiceId)
	}

	
	relevantService.ServiceDescription = utils.GetStringFromPointer(r.ServiceDescription)
	relevantService.ServiceDate = utils.ParseStringPointer(r.ServiceDate)
	relevantService.Firefighters = r.Firefighters
	relevantService.People = r.People
	relevantService.Infrastructures = r.Infrastructures
	relevantService.Antares = r.Antares
	relevantService.Vehicles = r.Vehicles
	relevantService.ServiceLocations = r.ServiceLocations
	relevantService.ServiceStations = r.ServiceStations
	relevantService.Centers = r.Centers
	relevantService.Unharmed = utils.ParseInt64StringPointer(r.Unharmed)
	relevantService.Injured = utils.ParseInt64StringPointer(r.Injured)
	relevantService.Transported = utils.ParseInt64StringPointer(r.Transported)

	relevantService.Deceased = utils.ParseInt64StringPointer(r.Deceased)

	relevantService.IsImportant = r.IsImportant

	if len(r.OprativeAreaName) == 0 {
		relevantService.OprativeAreaName = make([]string, 0)
	} else {
		relevantService.OprativeAreaName = r.OprativeAreaName
	}

	if len(r.Units) == 0 {
		relevantService.Units = make([]string, 0)
	} else {
		relevantService.Units = r.Units
	}

	return relevantService
}

func (r *RelevantServicesJson) ToModel() models.RelevantServices {
	relevantService := models.RelevantServices{}

	serviceId := utils.ParseInt64(r.ServiceId)

	relevantService.Id = utils.ParseInt64(r.Id) // Assuming utils.ConvertToPgTypeInt4 converts string to int32
	relevantService.RegionAreaDescription = &r.RegionAreaDescription
	relevantService.MissionCode = &r.MissionCode

	//operativeAreaName := make([]*string, len(r.OprativeAreaName))
	//units := make([]*string, len(r.Units))

	relevantService.OprativeAreaName = r.OprativeAreaName

	relevantService.ServiceId = &serviceId
	relevantService.ServiceDescription = &r.ServiceDescription
	relevantService.ServiceDate = &r.ServiceDate
	relevantService.Units = r.Units
	relevantService.Firefighters = r.Firefighters
	relevantService.People = r.People
	relevantService.Infrastructures = r.Infrastructures
	relevantService.Antares = r.Antares
	relevantService.Vehicles = r.Vehicles
	relevantService.ServiceLocations = r.ServiceLocations
	relevantService.ServiceStations = r.ServiceStations
	relevantService.Centers = r.Centers

	unharmed := utils.ParseInt64(r.Unharmed)
	injured := utils.ParseInt64(r.Injured)
	transported := utils.ParseInt64(r.Transported)
	deceased := utils.ParseInt64(r.Deceased)

	relevantService.Unharmed = &unharmed
	relevantService.Injured = &injured
	relevantService.Transported = &transported
	relevantService.Deceased = &deceased
	relevantService.IsImportant = r.IsImportant
	return relevantService
}
