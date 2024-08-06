package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionPersonJson struct {
	Id               string `json:"id"`
	ServiceId        string `json:"service_id"`
	UnitId           string `json:"unit_id"`
	InfrastructureId string `json:"infrastructure_id"`
	VehicleId        string `json:"vehicle_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Age              string `json:"age"`
	Gender           string `json:"gender"`
	LegalId          string `json:"legal_id"`
	Phone            string `json:"phone"`
	Employment       string `json:"employment"`
	State            string `json:"state"`
	Municipality     string `json:"municipality"`
	Parish           string `json:"parish"`
	Address          string `json:"address"`
	Pathology        string `json:"pathology"`
	Observations     string `json:"observations"`
	Condition        string `json:"condition"`
}

func ModelToMissionPersonJson(s models.MissionPerson) *MissionPersonJson {
	person := MissionPersonJson{}

	person.Id = utils.ConvertFromInt4(s.Id)
	person.UnitId = utils.ConvertFromInt4(s.UnitId)
	person.ServiceId = utils.ConvertFromInt4(s.ServiceId)

	person.InfrastructureId = utils.ConvertFromInt4(s.InfrastructureId)
	person.VehicleId = utils.ConvertFromInt4(s.VehicleId)
	person.FirstName = utils.ConvertFromText(s.FirstName)
	person.LastName = utils.ConvertFromText(s.LastName)
	person.Age = utils.ConvertFromInt2(s.Age)
	person.Gender = utils.ConvertFromText(s.Gender)
	person.LegalId = utils.ConvertFromText(s.LegalId)
	person.Phone = utils.ConvertFromText(s.Phone)
	person.Employment = utils.ConvertFromText(s.Employment)
	person.State = utils.ConvertFromText(s.State)
	person.Municipality = utils.ConvertFromText(s.Municipality)
	person.Address = utils.ConvertFromText(s.Address)
	person.Pathology = utils.ConvertFromText(s.Pathology)
	person.Observations = utils.ConvertFromText(s.Observations)
	person.Condition = utils.ConvertFromText(s.Condition)

	return &person
}

func (s *MissionPersonJson) ToModel() models.MissionPerson {
	person := models.MissionPerson{}

	person.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	person.UnitId = utils.ConvertToPgTypeInt4(utils.ParseInt(s.UnitId))
	person.ServiceId = utils.ConvertToPgTypeInt4(utils.ParseInt(s.ServiceId))

	person.InfrastructureId = utils.ConvertToPgTypeInt4(utils.ParseInt(s.InfrastructureId))
	person.VehicleId = utils.ConvertToPgTypeInt4(utils.ParseInt(s.VehicleId))
	person.FirstName = utils.ConvertToPgTypeText(s.FirstName)
	person.LastName = utils.ConvertToPgTypeText(s.LastName)
	person.Age = utils.ConvertToPgTypeInt2(utils.ParseInt(s.Age))
	person.Gender = utils.ConvertToPgTypeText(s.Gender)
	person.LegalId = utils.ConvertToPgTypeText(s.LegalId)
	person.Phone = utils.ConvertToPgTypeText(s.Phone)
	person.Employment = utils.ConvertToPgTypeText(s.Employment)
	person.State = utils.ConvertToPgTypeText(s.State)
	person.Municipality = utils.ConvertToPgTypeText(s.Municipality)
	person.Address = utils.ConvertToPgTypeText(s.Address)
	person.Pathology = utils.ConvertToPgTypeText(s.Pathology)
	person.Observations = utils.ConvertToPgTypeText(s.Observations)
	person.Condition = utils.ConvertToPgTypeText(s.Condition)

	return person
}
