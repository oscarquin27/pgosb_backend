package mission_infra_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionInfrastructureController struct {
	missionService services.MissionInfrastructureService
}

func NewMissionController(missionService services.MissionInfrastructureService) *MissionInfrastructureController {
	return &MissionInfrastructureController{
		missionService: missionService,
	}
}

func (u *MissionInfrastructureController) GetInfrastructure(c *gin.Context) {
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

	missionDto := api_models.ModelToMissionInfrastructureJson(*mission)

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionInfrastructureController) GetByServiceId(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))

	mission, err := u.missionService.GetByServiceId(id)

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionInfrastructureJson

	for _, s := range mission {
		newMission := api_models.ModelToMissionInfrastructureJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionInfrastructureController) GetAll(c *gin.Context) {

	mission, err := u.missionService.GetAll()

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionInfrastructureJson

	for _, s := range mission {
		newMission := api_models.ModelToMissionInfrastructureJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionInfrastructureController) Create(c *gin.Context) {
	var infra api_models.MissionInfrastructureJson

	if err := c.BindJSON(&infra); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	infrastructure := infra.ToModel()

	err := u.missionService.Create(&infrastructure)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Infraestructura creada satisfactoriamente")
}

func (u *MissionInfrastructureController) Update(c *gin.Context) {
	var infra api_models.MissionInfrastructureJson

	if err := c.BindJSON(&infra); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	infrastructure := infra.ToModel()

	err := u.missionService.Update(&infrastructure)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Infraestructura actualizada satisfactoriamente")
}

func (u *MissionInfrastructureController) Delete(c *gin.Context) {

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
	c.JSON(http.StatusOK, "Infraestructura eliminada satisfactoriamente")

}
