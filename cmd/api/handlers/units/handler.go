package units_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnitController struct {
	unityService services.UnitService
}

func NewUnityController(unityService services.UnitService) *UnitController {
	return &UnitController{
		unityService: unityService,
	}
}

func (u *UnitController) GetUnit(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	vehicle, err := u.unityService.Get(id)

	if err != nil {
		if err == models.ErrorUnitNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (u *UnitController) GetAllUnits(c *gin.Context) {

	vehicle, err := u.unityService.GetAll()

	if err != nil {
		if err == models.ErrorUnitNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (u *UnitController) CreateUnit(c *gin.Context) {
	var userDto api_models.UnitJson
	if err := c.BindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user := userDto.ToModel()

	err := u.unityService.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Unidad creada satisfactoriamente")
}

func (u *UnitController) UpdateUnit(c *gin.Context) {
	var userDto api_models.UnitJson

	if err := c.BindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user := userDto.ToModel()

	err := u.unityService.Update(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Unidad actualizada satisfactoriamente")
}

func (u *UnitController) DeleteUnit(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.unityService.Delete(id)

	if err != nil {
		if err == models.ErrorUnitNotUpdated {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Unidad eliminada satisfactoriamente")
}
