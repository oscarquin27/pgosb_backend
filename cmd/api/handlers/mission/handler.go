package mission_handlers

import (
	api_models "fdms/cmd/api/models"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionController struct {
	missionService services.MissionService
}

func NewMissionController(missionService services.MissionService) *MissionController {
	return &MissionController{
		missionService: missionService,
	}
}

func (u *MissionController) GetMission(c *gin.Context) {

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

	c.JSON(http.StatusOK, api_models.ModelToMissionJson(*mission))
}

func (u *MissionController) GetAllMissions(c *gin.Context) {

	mission, err := u.missionService.GetAllMissionSummary()

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionSummaryJson = make([]api_models.MissionSummaryJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionSummaryJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)

}

func (u *MissionController) Create(c *gin.Context) {
	var mission api_models.MissionJson

	if err := c.BindJSON(&mission); err != nil {
		logger.Error().Err(err).Msg("Error parseando modelo de mission json")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	missionEntity := mission.ToModel()

	id, err := u.missionService.Create(&missionEntity)

	if err != nil {
		logger.Error().Err(err).Msg("Error Creando Mission")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	missionEntity.Id = id.Id

	returnMissionValue := api_models.ModelToMissionJson(missionEntity)

	c.JSON(http.StatusOK, returnMissionValue)
}

func (u *MissionController) Update(c *gin.Context) {
	var missionJson api_models.MissionJson

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

	c.JSON(http.StatusOK, "Misión actualizada satisfactoriamente")
}

func (u *MissionController) Delete(c *gin.Context) {

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
	c.JSON(http.StatusOK, "Misión eliminado satisfactoriamente")

}
