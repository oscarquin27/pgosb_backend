package mission_service_handlers

import (
	api_models "fdms/cmd/api/models"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"fdms/src/utils/results"
	"net/http"
	"strings"

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

	id := utils.ParseInt64(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	mission, err := u.missionService.Get(id, isTemplate)

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

	id := utils.ParseInt64(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	mission, err := u.missionService.GetByMissionId(id, isTemplate)

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

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	mission, err := u.missionService.GetAll(isTemplate)

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

func (u *MissionServiceController) GetAllSummary(c *gin.Context) {

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	mission, err := u.missionService.GetAllMissionServiceSummary(isTemplate)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionServiceSummaryJson = make([]api_models.MissionServiceSummaryJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionServiceSummaryJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionServiceController) GetUnits(c *gin.Context) {

	id := utils.ParseInt64(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	result := u.missionService.GetUnits(id, isTemplate)

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

	id := utils.ParseInt64(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	result := u.missionService.GetUsers(id, isTemplate)

	if !result.IsSuccessful {

		logger.Warn().Err(result.Err.AssociateException()).
			Msg("Problemas ejecutando GetAllSimple")

		if result.Err.Code() == results.NotFoundErr {
			c.JSON(http.StatusOK, result.Value)
		}

		c.JSON(http.StatusInternalServerError, result.Value)
		return
	}

	var missionDto []api_models.MissionUserServiceJson = make([]api_models.MissionUserServiceJson, 0)

	for _, s := range result.Value {
		newMission := api_models.ModelToMissionUserServiceJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionServiceController) Create(c *gin.Context) {
	var missionJson api_models.MissionServiceJson

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

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

	id, err := u.missionService.Create(&mission, isTemplate)

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

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	err := u.missionService.Update(&mission, isTemplate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Servicio actualizado satisfactoriamente")

}

func (u *MissionServiceController) Delete(c *gin.Context) {

	id := utils.ParseInt64(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	err := u.missionService.Delete(id, isTemplate)

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

func (u *MissionServiceController) GetRelevantServices(c *gin.Context) {

	id := string(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	mission, err := u.missionService.GetRelevantServices(id, isTemplate)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.RelevantServicesJson = make([]api_models.RelevantServicesJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToRelevantServicesJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
	return
}

func (u *MissionServiceController) GetRelevantMissions(c *gin.Context) {

	id := string(c.Param("id"))

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	mission, err := u.missionService.GetRelevantMissions(id, isTemplate)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.RelevantServicesJson = make([]api_models.RelevantServicesJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToRelevantServicesJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
	return
}
