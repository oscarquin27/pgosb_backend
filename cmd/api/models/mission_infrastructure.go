package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type MissionInfrastructureJson struct {
	Id		        	string   `json:"id" binding:"required"`
	ServiceId      	    string   `json:"service_id"`
	BuildType       	string   `json:"build_type"`
	BuildOccupation     string   `json:"build_occupation"`
	BuildArea           string   `json:"build_area"`
	BuildAccess         string   `json:"build_access"`
	Levels              string   `json:"levels"`
	People              string   `json:"people"`
	GoodsType           string   `json:"goods_type"`
	BuildRoof           string   `json:"build_roof"`
	BuildFloor          string   `json:"build_floor"`
	BuildRoomType       string   `json:"build_room_type"`
	Observations        string   `json:"observations"`
}

func ModelToMissionInfrastructureJson(s models.MissionInfrastructure) *MissionInfrastructureJson {
	infra := MissionInfrastructureJson{}

	infra.Id = utils.ConvertFromInt4(s.Id)
	infra.ServiceId = utils.ConvertFromInt4(s.ServiceId)
	infra.BuildType = utils.ConvertFromText(s.BuildType)
	infra.BuildOccupation = utils.ConvertFromText(s.BuildOccupation)
	infra.BuildArea = utils.ConvertFromText(s.BuildArea)
	infra.BuildAccess = utils.ConvertFromText(s.BuildAccess)
	infra.Levels = utils.ConvertFromInt2(s.Levels)
	infra.People = utils.ConvertFromInt2(s.People)
	infra.GoodsType = utils.ConvertFromText(s.GoodsType)
	infra.BuildRoof = utils.ConvertFromText(s.BuildRoof)
	infra.BuildFloor = utils.ConvertFromText(s.BuildFloor)
	infra.BuildRoomType = utils.ConvertFromText(s.BuildRoomType)
	infra.Observations = utils.ConvertFromText(s.Observations)

	return &infra
}

func (s *MissionInfrastructureJson) ToModel() models.MissionInfrastructure {
	infra := models.MissionInfrastructure{}

	infra.Id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Id))
	infra.ServiceId = utils.ConvertToPgTypeInt4(utils.ParseInt(s.ServiceId))
	infra.BuildType = utils.ConvertToPgTypeText(s.BuildType)
	infra.BuildOccupation = utils.ConvertToPgTypeText(s.BuildOccupation)
	infra.BuildArea = utils.ConvertToPgTypeText(s.BuildArea)
	infra.BuildAccess = utils.ConvertToPgTypeText(s.BuildAccess)
	infra.Levels = utils.ConvertToPgTypeInt2(utils.ParseInt(s.Levels))
	infra.People = utils.ConvertToPgTypeInt2(utils.ParseInt(s.People))
	infra.GoodsType = utils.ConvertToPgTypeText(s.GoodsType)
	infra.BuildRoof = utils.ConvertToPgTypeText(s.BuildRoof)
	infra.BuildFloor = utils.ConvertToPgTypeText(s.BuildFloor)
	infra.BuildRoomType = utils.ConvertToPgTypeText(s.BuildRoomType)
	infra.Observations = utils.ConvertToPgTypeText(s.Observations)
	
	return infra
}
