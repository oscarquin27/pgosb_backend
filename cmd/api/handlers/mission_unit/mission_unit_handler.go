package mission_unit_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionUnitController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionUnit, api_models.MissionUnitJson]
	missionUnitService     services.MissionUnitService
}

func NewMissionUnitController(missionUnitService services.MissionUnitService) *MissionUnitController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionUnit, api_models.MissionUnitJson](missionUnitService)

	return &MissionUnitController{
		abstractServiceHandler: *abstractHandler,
		missionUnitService:     missionUnitService,
	}
}

func (u *MissionUnitController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionUnitJson, c)
}

func (u *MissionUnitController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToMissionUnitJson, c)
}

func (u *MissionUnitController) Create(c *gin.Context) {
	s := api_models.MissionUnitJson{}

	var model abstract_handler.AbstactModel[models.MissionUnit, api_models.MissionUnitJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionUnitJson, c)
}

func (u *MissionUnitController) Update(c *gin.Context) {

	s := api_models.MissionUnitJson{}

	var model abstract_handler.AbstactModel[models.MissionUnit, api_models.MissionUnitJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionUnitJson, c)
}

func (u *MissionUnitController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionUnitController) GetAllSummary(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	result, err := u.missionUnitService.GetByMissionId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	firefighters := []api_models.MissionUnitSummaryJson{}

	for _, firefighter := range result {
		firefighters = append(firefighters, *api_models.ModelToMissionUnitSummaryJson(&firefighter))
	}

	c.JSON(http.StatusOK, firefighters)
}
