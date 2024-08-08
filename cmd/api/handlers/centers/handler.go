package center_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CenterController struct {
	centerService services.CentersService
}

func NewCenterController(centerService services.CentersService) *CenterController {
	return &CenterController{
		centerService: centerService,
	}
}

func (u *CenterController) GetCenter(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	center, err := u.centerService.Get(id)

	if err != nil {
		if err == models.ErrorCenterFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	centerResult := api_models.ModelToCentersJson(*center)

	c.JSON(http.StatusOK, centerResult)
}

func (u *CenterController) GetAllCenters(c *gin.Context) {

	centers, err := u.centerService.GetAll()

	if err != nil {
		if err == models.ErrorCenterFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var centersResult []api_models.CentersJson

	for _, c := range centers {
		newCenter := api_models.ModelToCentersJson(c)
		centersResult = append(centersResult, *newCenter)
	}

	c.JSON(http.StatusOK, centersResult)

}

func (u *CenterController) Create(c *gin.Context) {
	var centerDto api_models.CentersJson

	if err := c.BindJSON(&centerDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	center := centerDto.ToModel()

	err := u.centerService.Create(&center)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "centro asistencial creado satisfactoriamente")
}

func (u *CenterController) Update(c *gin.Context) {
	var centerDto api_models.CentersJson

	if err := c.BindJSON(&centerDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	center := centerDto.ToModel()

	err := u.centerService.Update(&center)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Centro asistencial actualizado satisfactoriamente")
}

func (u *CenterController) Delete(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	err := u.centerService.Delete(id)

	if err != nil {
		if err == models.ErrorCenterNotUpdated {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Centro asistencial borrado satisfactoriamente")
}
