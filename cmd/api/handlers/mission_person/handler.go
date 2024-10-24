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

	missionDto := api_models.ModelToMissionPersonJson(*mission)

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionPersonController) GetMissionId(c *gin.Context) {

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

	var missionDto []api_models.MissionPersonJson = make([]api_models.MissionPersonJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionPersonJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

func (u *MissionPersonController) GetAll(c *gin.Context) {

	mission, err := u.missionService.GetAll()

	if err != nil {
		if err == models.ErrorMissionNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var missionDto []api_models.MissionPersonJson = make([]api_models.MissionPersonJson, 0)

	for _, s := range mission {
		newMission := api_models.ModelToMissionPersonJson(s)
		missionDto = append(missionDto, *newMission)
	}

	c.JSON(http.StatusOK, missionDto)
}

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
	var person api_models.MissionPersonJson

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	p := person.ToModel()

	err := u.missionService.Create(&p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Persona creada satisfactoriamente")
}

func (u *MissionPersonController) Update(c *gin.Context) {
	var person api_models.MissionPersonJson

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	p := person.ToModel()

	err := u.missionService.Update(&p)

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
