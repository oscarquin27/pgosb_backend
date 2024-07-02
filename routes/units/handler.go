package routes

import (
	entities "fdms/domain/entities/units"
	unity "fdms/domain/units"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnitController struct {
	unityService unity.UnitRepository
}

func NewUnityController(unityService unity.UnitRepository) *UnitController {
	return &UnitController{
		unityService : unityService,
	}
}

func (u *UnitController) GetUnity(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	vehicle, err := u.unityService.GetUnit(id)

	if err != nil {
		if err == entities.ErrorUnitNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
	return
}

func (u *UnitController) GetAllUnities(c *gin.Context){

	vehicle, err := u.unityService.GetAll()

	if err != nil {
		if err == entities.ErrorUnitNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
	return
}

func (u *UnitController) CreateUnity(c *gin.Context){
	var user entities.Unit
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.unityService.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Unidad creada satisfactoriamente")
}


func (u *UnitController) UpdateUnity(c *gin.Context){
	var user entities.Unit

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.unityService.Update(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Unidad actualizada satisfactoriamente")
}

func (u *UnitController) DeleteUnity(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.unityService.Delete(id)

	if err != nil {
		if err == entities.ErrorUnitNotUpdated {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Unidad eliminada satisfactoriamente")
	return
}
