package mission_person_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissionPersonController struct {
	missionService services.MissionPersonService
}

func NewMissionPersonController(missionService services.MissionPersonService) *MissionPersonController {
	return &MissionPersonController{
		missionService: missionService,
	}
}

func (u *MissionPersonController) Get(c *gin.Context) {

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

	var missionDto []api_models.MissionPersonJson

	for _, s := range mission {
		newMission := api_models.ModelToMissionPersonJson(s)
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

func (u *MissionPersonController) Create(c *gin.Context) {
	var mission models.MissionPerson

	if err := c.BindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.missionService.Create(&mission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Persona creada satisfactoriamente")
}

func (u *MissionPersonController) Update(c *gin.Context) {
	var mission models.MissionPerson

	if err := c.BindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.missionService.Update(&mission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Persona actualizada satisfactoriamente")
}

func (u *MissionPersonController) Delete(c *gin.Context) {

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
	c.JSON(http.StatusOK, "Persona eliminada satisfactoriamente")

}