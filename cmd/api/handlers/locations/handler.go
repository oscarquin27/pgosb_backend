package location_handlers

import (
	api_models "fdms/cmd/api/models"
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

func (u *LocationController) GetState(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	states, err := u.locationService.GetState(id)

	if err != nil {
		if err == models.ErrorStateFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, states)
}

func (u *LocationController) GetAllStates(c *gin.Context) {

	states, err := u.locationService.GetAllStates()

	if err != nil {
		if err == models.ErrorStateFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, states)

}

func (u *LocationController) CreateState(c *gin.Context) {
	var location models.State
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.CreateState(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Estado creado satisfactoriamente")
}

func (u *LocationController) UpdateState(c *gin.Context) {
	var location models.State

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.UpdateState(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Estado actualizado satisfactoriamente")
}

func (u *LocationController) DeleteState(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteState(id)

	if err != nil {
		if err == models.ErrorStateNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Estado eliminado satisfactoriamente")
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

func (u *LocationController) GetMunicipality(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.locationService.GetMunicipality(id)

	if err != nil {
		if err == models.ErrorMunicipalityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}

func (u *LocationController) GetAllMunicipality(c *gin.Context) {

	user, err := u.locationService.GetAllMunicipality()

	if err != nil {
		if err == models.ErrorMunicipalityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}

func (u *LocationController) CreateMunicipality(c *gin.Context) {
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

	c.JSON(http.StatusCreated, "Ciudad creada satisfactoriamente")
}

func (u *LocationController) UpdateMunicipality(c *gin.Context) {
	var location models.Municipality

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.locationService.UpdateMunicipality(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Municipio actualizado satisfactoriamente")
}

func (u *LocationController) DeleteMunicipality(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteMunicipality(id)

	if err != nil {
		if err == models.ErrorCityNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Municipio eliminada satisfactoriamente")

}

func (u *LocationController) GetParish(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.locationService.GetParish(id)

	if err != nil {
		if err == models.ErrorParishFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}

func (u *LocationController) GetAllParish(c *gin.Context) {

	user, err := u.locationService.GetAllParish()

	if err != nil {
		if err == models.ErrorParishFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}

func (u *LocationController) CreateParish(c *gin.Context) {
	var location models.Parish
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.locationService.CreateParish(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Ciudad creada satisfactoriamente")
}

func (u *LocationController) UpdateParish(c *gin.Context) {
	var location models.Parish

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.locationService.UpdateParish(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Parroquia actualizada satisfactoriamente")
}

func (u *LocationController) DeleteParish(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteParish(id)

	if err != nil {
		if err == models.ErrorCityNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Parroquia eliminada satisfactoriamente")

}

func (u *LocationController) GetStation(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	fire_station, err := u.locationService.GetStation(id)

	if err != nil {
		if err == models.ErrorStationFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fire_station_json := api_models.ModelToStationJson(*fire_station)
	c.JSON(http.StatusOK, fire_station_json)

}

func (u *LocationController) GetAllStations(c *gin.Context) {

	fire_station, err := u.locationService.GetAllStations()

	if err != nil {
		if err == models.ErrorParishFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fireStations := []api_models.StationJson{}

	for _, st := range fire_station {
		newStation := api_models.ModelToStationJson(st)
		fireStations = append(fireStations, *newStation)
	}

	c.JSON(http.StatusOK, fireStations)

}

func (u *LocationController) CreateStation(c *gin.Context) {
	var locationDto api_models.StationJson

	if err := c.BindJSON(&locationDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	location := locationDto.ToModel()

	err := u.locationService.CreateStation(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Estación creada satisfactoriamente")
}

func (u *LocationController) UpdateStation(c *gin.Context) {
	var locationDto api_models.StationJson

	if err := c.BindJSON(&locationDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	location := locationDto.ToModel()

	err := u.locationService.UpdateStation(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Estación actualizada satisfactoriamente")
}

func (u *LocationController) DeleteStation(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteStation(id)

	if err != nil {
		if err == models.ErrorStationNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Estación eliminada satisfactoriamente")
}
