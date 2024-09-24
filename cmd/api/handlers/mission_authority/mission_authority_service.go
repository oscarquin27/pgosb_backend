package mission_authority_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MissionAuthorityServiceController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionAuthorityService, api_models.MissionAuthorityServiceJson]
	authorityService       services.MissionAuthorityRelateService
}

func NewMissionAuthorityServiceController(service abstract_handler.AbstractCRUDService[models.MissionAuthorityService],
	authorityService services.MissionAuthorityRelateService) *MissionAuthorityServiceController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionAuthorityService, api_models.MissionAuthorityServiceJson](service)

	return &MissionAuthorityServiceController{
		abstractServiceHandler: *abstractHandler,
		authorityService:       authorityService,
	}
}

func (u *MissionAuthorityServiceController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionAuthorityServiceJson, c)
}

func (u *MissionAuthorityServiceController) GetAll(c *gin.Context) {
	u.abstractServiceHandler.GetAll(api_models.ModelToMissionAuthorityServiceJson, c)
}

func (u *MissionAuthorityServiceController) Create(c *gin.Context) {
	s := api_models.MissionAuthorityServiceJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthorityService, api_models.MissionAuthorityServiceJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionAuthorityServiceJson, c)
}

func (u *MissionAuthorityServiceController) Update(c *gin.Context) {
	s := api_models.MissionAuthorityServiceJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthorityService, api_models.MissionAuthorityServiceJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionAuthorityServiceJson, c)
}

func (u *MissionAuthorityServiceController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionAuthorityServiceController) GetByServiceId(c *gin.Context) {

	missionId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := u.authorityService.GetByServiceId(missionId)

	if !result.IsSuccessful {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Err.Message()})
		return
	}

	var data []api_models.MissionAuthorityServiceSummaryJson = make([]api_models.MissionAuthorityServiceSummaryJson, 0)

	for _, v := range result.Value {
		data = append(data, *api_models.ModelToMissionAuthorityServiceSummaryJson(&v))
	}

	c.JSON(http.StatusOK, data)

}
