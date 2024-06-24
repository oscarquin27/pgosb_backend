package routes

import (
	entities "fdms/domain/entities/unities"
	unity "fdms/domain/unities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnityController struct {
	unityService unity.UnityRepository
}

func NewUnityController(unityService unity.UnityRepository) *UnityController {
	return &UnityController{
		unityService : unityService,
	}
}

func (u *UnityController) GetUnity(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	vehicle, err := u.unityService.GetUnity(id)

	if err != nil {
		if err == entities.ErrorUnityNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
	return
}

func (u *UnityController) GetAllUnities(c *gin.Context){

	vehicle, err := u.unityService.GetAll()

	if err != nil {
		if err == entities.ErrorUnityNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
	return
}

func (u *UnityController) CreateUnity(c *gin.Context){
	var user entities.Unity
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


func (u *UnityController) UpdateUnity(c *gin.Context){
	var user entities.Unity

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

func (u *UnityController) DeleteUnity(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.unityService.Delete(id)

	if err != nil {
		if err == entities.ErrorUnityNotUpdated {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Unidad eliminada satisfactoriamente")
	return
}
