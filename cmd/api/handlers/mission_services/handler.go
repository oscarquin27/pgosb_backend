package mission_service_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
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

	var missionDto []api_models.MissionServiceJson

	for _, s := range mission {
		newMission := api_models.ModelToMissionServiceJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)}

// func (u *MissionServiceController) GetAllServices(c *gin.Context) {

// 	mission, err := u.missionService.GetAll()

// 	if err != nil {
// 		if err == models.ErrorMissionNotFound {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	var missionDto []api_models.MissionServiceJson

// 	for _, s := range mission {
// 		newMission := api_models.ModelToMissionServiceJson(s)
// 		missionDto = append(missionDto, *newMission)
// 	}

// 	c.JSON(http.StatusOK, missionDto)

// }

func (u *MissionServiceController) Create(c *gin.Context) {
	var mission models.MissionService

	if err := c.BindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

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

	c.JSON(http.StatusOK, id)
}

func (u *MissionServiceController) Update(c *gin.Context) {
	var mission models.MissionService

	if err := c.BindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

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