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

type MissionAuthorityController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionAuthority, api_models.MissionAuthorityJson]
	authorityService       services.MissionAuthorityService
}

func NewMissionAuthorityController(service abstract_handler.AbstractCRUDService[models.MissionAuthority],
	authorityService services.MissionAuthorityService) *MissionAuthorityController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionAuthority, api_models.MissionAuthorityJson](service)

	return &MissionAuthorityController{
		abstractServiceHandler: *abstractHandler,
		authorityService:       authorityService,
	}
}

func (u *MissionAuthorityController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionAuthorityJson, c)
}

func (u *MissionAuthorityController) GetAll(c *gin.Context) {
	u.abstractServiceHandler.GetAll(api_models.ModelToMissionAuthorityJson, c)
}

func (u *MissionAuthorityController) Create(c *gin.Context) {
	s := api_models.MissionAuthorityJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthority, api_models.MissionAuthorityJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionAuthorityJson, c)
}

func (u *MissionAuthorityController) Update(c *gin.Context) {
	s := api_models.MissionAuthorityJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthority, api_models.MissionAuthorityJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionAuthorityJson, c)
}

func (u *MissionAuthorityController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionAuthorityController) GetByMissionId(c *gin.Context) {

	missionId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := u.authorityService.GetByMissionId(missionId)

	if !result.IsSuccessful {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result.Value})
}

func (u *MissionAuthorityController) GetSummaryByMissionId(c *gin.Context) {

	missionId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := u.authorityService.GetSummaryByMissionId(missionId)

	if !result.IsSuccessful {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Err.Message()})
		return
	}

	var data []api_models.MissionAuthoritySummaryJson = make([]api_models.MissionAuthoritySummaryJson, 0)

	for _, v := range result.Value {
		data = append(data, *api_models.ModelToMissionAuthoritySummaryJson(&v))
	}

	c.JSON(http.StatusOK, data)

}
