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

type MissionAuthorityVehicleController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionAuthorityVehicle, api_models.MissionAuthorityVehicleJson]
	authorityService       services.MissionAuthorityVehicleService
}

func NewMissionAuthorityVehicleController(service abstract_handler.AbstractCRUDService[models.MissionAuthorityVehicle],
	authorityService services.MissionAuthorityVehicleService) *MissionAuthorityVehicleController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionAuthorityVehicle, api_models.MissionAuthorityVehicleJson](service)

	return &MissionAuthorityVehicleController{
		abstractServiceHandler: *abstractHandler,
		authorityService:       authorityService,
	}
}

func (u *MissionAuthorityVehicleController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionAuthorityVehicleJson, c)
}

func (u *MissionAuthorityVehicleController) GetAll(c *gin.Context) {
	u.abstractServiceHandler.GetAll(api_models.ModelToMissionAuthorityVehicleJson, c)
}

func (u *MissionAuthorityVehicleController) Create(c *gin.Context) {
	s := api_models.MissionAuthorityVehicleJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthorityVehicle, api_models.MissionAuthorityVehicleJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionAuthorityVehicleJson, c)
}

func (u *MissionAuthorityVehicleController) Update(c *gin.Context) {
	s := api_models.MissionAuthorityVehicleJson{}

	var model abstract_handler.AbstactModel[models.MissionAuthorityVehicle, api_models.MissionAuthorityVehicleJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionAuthorityVehicleJson, c)
}

func (u *MissionAuthorityVehicleController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionAuthorityVehicleController) GetByAuthorityId(c *gin.Context) {

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

	var data []api_models.MissionAuthorityVehicleJson = make([]api_models.MissionAuthorityVehicleJson, 0)

	for _, v := range result.Value {
		data = append(data, *api_models.ModelToMissionAuthorityVehicleJson(&v))
	}

	c.JSON(http.StatusOK, data)

}
