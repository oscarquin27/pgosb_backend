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
	v := api_models.ModelToUnitJson(vehicle)

	c.JSON(http.StatusOK, v)
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

	vehicleDto := []api_models.UnitJson{}

	for _, us := range vehicle {
		newVehicle := api_models.ModelToUnitJson(&us)
		vehicleDto = append(vehicleDto, *newVehicle)
	}

	c.JSON(http.StatusOK, vehicleDto)
}

func (u *UnitController) CreateUnit(c *gin.Context) {
	var unitDto api_models.UnitJson

	if err := c.BindJSON(&unitDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	unit := unitDto.ToModel()

	err := u.unityService.Create(&unit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Unidad creada satisfactoriamente")
}

func (u *UnitController) UpdateUnit(c *gin.Context) {
	var unitDto api_models.UnitJson

	if err := c.BindJSON(&unitDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	unit := unitDto.ToModel()

	err := u.unityService.Update(&unit)

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
