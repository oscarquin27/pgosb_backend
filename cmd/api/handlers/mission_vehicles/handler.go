package mission_vehicle_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionVehicleController struct {
	missionService services.MissionVehicleService
}

func NewMissionVehicleController(missionService services.MissionVehicleService) *MissionVehicleController {
	return &MissionVehicleController{
		missionService: missionService,
	}
}

func (u *MissionVehicleController) GetVehicle(c *gin.Context) {

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

	missionDto := api_models.ModelToMissionVehicleJson(*mission)

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionVehicleController) GetAll(c *gin.Context) {

	mission, err := u.missionService.GetAll()

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionVehicleJson = make([]api_models.MissionVehicleJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionVehicleJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionVehicleController) GetMissionId(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	mission, err := u.missionService.GetMissionId(id)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionVehicleJson = make([]api_models.MissionVehicleJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionVehicleJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionVehicleController) Create(c *gin.Context) {
	var vehicle api_models.MissionVehicleJson

	if err := c.BindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	vehicles := vehicle.ToModel()

	err := u.missionService.Create(&vehicles)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Servicio creado satisfactoriamente")
}

func (u *MissionVehicleController) Update(c *gin.Context) {
	var vehicle api_models.MissionVehicleJson

	if err := c.BindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	vehicles := vehicle.ToModel()

	err := u.missionService.Update(&vehicles)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Servicio actualizado satisfactoriamente")
}

func (u *MissionVehicleController) Delete(c *gin.Context) {

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
