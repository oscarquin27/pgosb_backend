package mission_location_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MissionLocationController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionLocation, api_models.MissionLocationJson]
	missionService         services.MissionLocationService
}

func NewMissionLocationController(stationService abstract_handler.AbstractCRUDService[models.MissionLocation],
	missionService services.MissionLocationService) *MissionLocationController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionLocation, api_models.MissionLocationJson](stationService)

	return &MissionLocationController{
		abstractServiceHandler: *abstractHandler,
		missionService:         missionService,
	}
}

func (u *MissionLocationController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMissionLocationJson, c)
}

func (u *MissionLocationController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToMissionLocationJson, c)
}

func (u *MissionLocationController) Create(c *gin.Context) {
	s := api_models.MissionLocationJson{}

	var model abstract_handler.AbstactModel[models.MissionLocation, api_models.MissionLocationJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMissionLocationJson, c)
}

func (u *MissionLocationController) Update(c *gin.Context) {

	s := api_models.MissionLocationJson{}

	var model abstract_handler.AbstactModel[models.MissionLocation, api_models.MissionLocationJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMissionLocationJson, c)
}

func (u *MissionLocationController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *MissionLocationController) GetLocationsByServiceId(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var returnList []api_models.MissionLocationJson = make([]api_models.MissionLocationJson, 0)

	r := u.missionService.GetLocationsByServiceId(id)

	if !r.IsSuccessful {

		logger.Warn().Err(r.Err.AssociateException()).
			Msgf("El get con Id:%d no fue exitoso", id)

		if r.Err.Code() == results.NotFoundErr {

			c.JSON(http.StatusOK, returnList)
			return

		}

		c.JSON(http.StatusInternalServerError, r.Err.Message())
		return

	}

	for _, val := range r.Value {
		newLocation := api_models.ModelToMissionLocationJson(&val)
		returnList = append(returnList, *newLocation)
	}

	c.JSON(http.StatusOK, returnList)
}
