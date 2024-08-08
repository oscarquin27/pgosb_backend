package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

type CentersJson struct {
	Id           string `json:"id" db:"id"`
	UrbId		 string `json:"urb_id"`
	Type    	 string `json:"type"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
}

func ModelToCentersJson(s models.Center) *CentersJson {
	center := &CentersJson{}

	center.Id = utils.ConvertIntToString(*s.Id)
	center.UrbId = utils.ConvertIntToString(*s.UrbId)
	center.Type = *s.Type
	center.Name = *s.Name
	center.Address = *s.Address
	center.Phone = *s.Phone

	return center
}

func (s *CentersJson) ToModel() models.Center {
	center := models.Center{}

	id := utils.ParseInt(s.Id)
	urbId := utils.ParseInt(s.UrbId)

	center.Id = &id
	center.UrbId = &urbId
	center.Type = &s.Type
	center.Name = &s.Name
	center.Address = &s.Address
	center.Phone = &s.Phone

	return center
}
