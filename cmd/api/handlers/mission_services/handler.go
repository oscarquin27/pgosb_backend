package mission_service_handlers

import (
	api_models "fdms/cmd/api/models"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"fdms/src/utils/results"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionServiceController struct {
	missionService services.MissionServiceService
}

func NewServiceServiceController(missionService services.MissionServiceService) *MissionServiceController {
	return &MissionServiceController{
		missionService: missionService,
	}
}

func (u *MissionServiceController) Get(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	mission, err := u.missionService.Get(id)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	missionDto := api_models.ModelToMissionServiceJson(*mission)

	c.JSON(http.StatusOK, missionDto)
	return
}

func (u *MissionServiceController) GetByMissionId(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	mission, err := u.missionService.GetByMissionId(id)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionServiceJson = make([]api_models.MissionServiceJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionServiceJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionServiceController) GetAll(c *gin.Context) {

	mission, err := u.missionService.GetAll()

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionServiceJson = make([]api_models.MissionServiceJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionServiceJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
	return
}

func (u *MissionServiceController) GetUnits(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	result := u.missionService.GetUnits(id)

	if !result.IsSuccessful {

		logger.Warn().Err(result.Err.AssociateException()).
			Msg("Problemas ejecutando GetAllSimple")

		if result.Err.Code() == results.NotFoundErr {
			c.JSON(http.StatusOK, result.Value)
		}

		c.JSON(http.StatusInternalServerError, result.Value)
		return
	}

	c.JSON(http.StatusOK, result.Value)
}

func (u *MissionServiceController) GetUsers(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	result := u.missionService.GetUsers(id)

	if !result.IsSuccessful {

		logger.Warn().Err(result.Err.AssociateException()).
			Msg("Problemas ejecutando GetAllSimple")

		if result.Err.Code() == results.NotFoundErr {
			c.JSON(http.StatusOK, result.Value)
		}

		c.JSON(http.StatusInternalServerError, result.Value)
		return
	}

	c.JSON(http.StatusOK, result.Value)
}

func (u *MissionServiceController) Create(c *gin.Context) {
	var missionJson api_models.MissionServiceJson

	if err := c.BindJSON(&missionJson); err != nil {

		logger.Error().Err(err).Msg("Error Parseando MissionService")

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mission := missionJson.ToModel()

	mission.Id.Valid = true
	mission.AntaresId.Valid = true
	mission.MissionId.Valid = true
	mission.Summary.Valid = true
	mission.Description.Valid = true

	id, err := u.missionService.Create(&mission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	missionJson = *api_models.ModelToMissionServiceJson(*id)

	c.JSON(http.StatusOK, missionJson)

}

func (u *MissionServiceController) Update(c *gin.Context) {
	var missionJson api_models.MissionServiceJson

	if err := c.BindJSON(&missionJson); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mission := missionJson.ToModel()

	err := u.missionService.Update(&mission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Servicio actualizado satisfactoriamente")

}

func (u *MissionServiceController) Delete(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	err := u.missionService.Delete(id)

	if err != nil {
		if err == models.ErrorUserNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Servicio eliminado satisfactoriamente")

}
