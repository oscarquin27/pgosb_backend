package mission_firefighter_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionFireFigtherController struct {
	abstractServiceHandler    abstract_handler.AbstractHandler[models.MissionFirefighter, api_models.MissionFirefighterJson]
	missionFirefighterService services.MissionFirefighterService
}

func NewMissionFireFigtherController(stationService abstract_handler.AbstractCRUDService[models.MissionFirefighter], missionFirefighterService services.MissionFirefighterService) *MissionFireFigtherController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionFirefighter, api_models.MissionFirefighterJson](stationService)

	return &MissionFireFigtherController{
		abstractServiceHandler:    *abstractHandler,
		missionFirefighterService: missionFirefighterService,
	}
}

func (u *MissionFireFigtherController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionFirefighterJson, c)
}

func (u *MissionFireFigtherController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToMissionFirefighterJson, c)
}

func (u *MissionFireFigtherController) Create(c *gin.Context) {
	s := api_models.MissionFirefighterJson{}

	var model abstract_handler.AbstactModel[models.MissionFirefighter, api_models.MissionFirefighterJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionFirefighterJson, c)
}

func (u *MissionFireFigtherController) Update(c *gin.Context) {

	s := api_models.MissionFirefighterJson{}

	var model abstract_handler.AbstactModel[models.MissionFirefighter, api_models.MissionFirefighterJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionFirefighterJson, c)
}

func (u *MissionFireFigtherController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionFireFigtherController) GetUsers(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	result, err := u.missionFirefighterService.GetByMissionId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	firefighters := []api_models.MissionFirefighterUserJson{}

	for _, firefighter := range result {
		firefighters = append(firefighters, *api_models.ModelToMissionFirefighterUserJson(&firefighter))
	}

	c.JSON(http.StatusOK, firefighters)
}
