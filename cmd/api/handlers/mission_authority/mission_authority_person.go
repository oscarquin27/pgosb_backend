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

type MissionAuthorityPersonController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionAuthorityPerson, api_models.MissionAuthorityPersonJson]
	authorityService       services.MissionAuthorityPersonService
}

func NewMissionAuthorityPersonController(service abstract_handler.AbstractCRUDService[models.MissionAuthorityPerson],
	authorityService services.MissionAuthorityPersonService) *MissionAuthorityPersonController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionAuthorityPerson, api_models.MissionAuthorityPersonJson](service)

	return &MissionAuthorityPersonController{
		abstractServiceHandler: *abstractHandler,
		authorityService:       authorityService,
	}
}

func (u *MissionAuthorityPersonController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionAuthorityPersonJson, c)
}

func (u *MissionAuthorityPersonController) GetAll(c *gin.Context) {
	u.abstractServiceHandler.GetAll(api_models.ModelToMissionAuthorityPersonJson, c)
}

func (u *MissionAuthorityPersonController) Create(c *gin.Context) {
	s := api_models.MissionAuthorityPersonJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthorityPerson, api_models.MissionAuthorityPersonJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionAuthorityPersonJson, c)
}

func (u *MissionAuthorityPersonController) Update(c *gin.Context) {
	s := api_models.MissionAuthorityPersonJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthorityPerson, api_models.MissionAuthorityPersonJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionAuthorityPersonJson, c)
}

func (u *MissionAuthorityPersonController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionAuthorityPersonController) GetByAuthorityId(c *gin.Context) {

	missionId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := u.authorityService.GetByAuthorityId(missionId)

	if !result.IsSuccessful {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Err.Message()})
		return
	}

	var data []api_models.MissionAuthorityPersonJson = make([]api_models.MissionAuthorityPersonJson, 0)

	for _, v := range result.Value {
		data = append(data, *api_models.ModelToMissionAuthorityPersonJson(&v))
	}

	c.JSON(http.StatusOK, data)

}
