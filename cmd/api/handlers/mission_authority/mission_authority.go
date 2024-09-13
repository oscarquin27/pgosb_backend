package mission_authority_handler

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

type MissionAuthorityController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.MissionAuthority, api_models.MissionAuthorityJson]
	missionService         services.MissionAuthorityService
}

func NewMissionAuthorityController(missionService services.MissionAuthorityService) *MissionAuthorityController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.MissionAuthority, api_models.MissionAuthorityJson](missionService)

	return &MissionAuthorityController{
		abstractServiceHandler: *abstractHandler,
		missionService:         missionService,
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

func (u *MissionAuthorityController) GetLocationsByServiceId(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var returnList []api_models.MissionAuthorityJson = make([]api_models.MissionAuthorityJson, 0)

	r := u.missionService.GetByServiceId(id)

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
		newLocation := api_models.ModelToMissionAuthorityJson(&val)
		returnList = append(returnList, *newLocation)
	}

	c.JSON(http.StatusOK, returnList)
}
