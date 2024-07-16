package mission_handlers

import (
	api_models "fdms/cmd/api/models"
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

	mission, err := u.missionService.GetAll()

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionJson

	for _, s := range mission {
		newMission := api_models.ModelToMissionJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)

}

func (u *MissionController) Create(c *gin.Context) {
	var mission models.Mission

	if err := c.BindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.missionService.Create(&mission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Misión creada satisfactoriamente")
}

func (u *MissionController) Update(c *gin.Context) {
	var mission models.Mission

	if err := c.BindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

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