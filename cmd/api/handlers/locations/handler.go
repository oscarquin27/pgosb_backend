package location_handlers

import (
	"fdms/src/models"
	"fdms/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	locationService services.LocationsService
}

func NewLocationController(locationService services.LocationsService) *LocationController {
	return &LocationController{
		locationService: locationService,
	}
}

func (u *LocationController) GetCity(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.locationService.GetCity(id)

	if err != nil {
		if err == models.ErrorCityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *LocationController) GetAllCity(c *gin.Context) {

	cities, err := u.locationService.GetAllCity()

	if err != nil {
		if err == models.ErrorCityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cities)
}

func (u *LocationController) CreateCity(c *gin.Context) {
	var location models.City
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.locationService.CreateCity(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Ciudad creada satisfactoriamente")
}

func (u *LocationController) UpdateCity(c *gin.Context) {
	var location models.City

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.locationService.UpdateCity(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Estado actualizado satisfactoriamente")
}

func (u *LocationController) DeleteCity(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteCity(id)

	if err != nil {
		if err == models.ErrorCityNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Ciudad eliminada satisfactoriamente")

}
