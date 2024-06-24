package routes

import (
	city "fdms/domain/entities/cities"
	municipality "fdms/domain/entities/municipalities"
	parish "fdms/domain/entities/parish"
	state "fdms/domain/entities/states"
	station "fdms/domain/entities/stations"
	locations "fdms/domain/locations"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	locationService locations.LocationsRepository
}

func NewLocationController(locationService locations.LocationsRepository) *LocationController {
	return &LocationController{
		locationService : locationService,
	}
}

func (u *LocationController) GetState(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	states, err := u.locationService.GetState(id)

	if err != nil {
		if err == state.ErrorStateFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, states)
	return
}

func (u *LocationController) GetAllStates(c *gin.Context){

	states, err := u.locationService.GetAllStates()

	if err != nil {
		if err == state.ErrorStateFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, states)
	return
}

func (u *LocationController) CreateState(c *gin.Context){
	var location state.State
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

func (u *LocationController) UpdateState(c *gin.Context){
	var location state.State

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

func (u *LocationController) DeleteState(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteState(id)

	if err != nil {
		if err == state.ErrorStateNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Estado eliminado satisfactoriamente")
	return
}

func (u *LocationController) GetCity(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.locationService.GetCity(id)

	if err != nil {
		if err == city.ErrorCityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *LocationController) GetAllCity(c *gin.Context){

	cities , err := u.locationService.GetAllCity()

	if err != nil {
		if err == city.ErrorCityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cities)
	return
}

func (u *LocationController) CreateCity(c *gin.Context){
	var location city.City
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.CreateCity(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Ciudad creada satisfactoriamente")
}

func (u *LocationController) UpdateCity(c *gin.Context){
	var location city.City

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.UpdateCity(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Estado actualizado satisfactoriamente")
}

func (u *LocationController) DeleteCity(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteCity(id)

	if err != nil {
		if err == city.ErrorCityNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Ciudad eliminada satisfactoriamente")
	return
}

func (u *LocationController) GetMunicipality(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.locationService.GetMunicipality(id)

	if err != nil {
		if err == municipality.ErrorMunicipalityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *LocationController) GetAllMunicipality(c *gin.Context){

	user, err := u.locationService.GetAllMunicipality()

	if err != nil {
		if err == municipality.ErrorMunicipalityFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *LocationController) CreateMunicipality(c *gin.Context){
	var location city.City
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.CreateCity(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusCreated, "Ciudad creada satisfactoriamente")
}

func (u *LocationController) UpdateMunicipality(c *gin.Context){
	var location municipality.Municipality

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.UpdateMunicipality(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Municipio actualizado satisfactoriamente")
}

func (u *LocationController) DeleteMunicipality(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteMunicipality(id)

	if err != nil {
		if err == city.ErrorCityNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Municipio eliminada satisfactoriamente")
	return
}

func (u *LocationController) GetParish(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.locationService.GetParish(id)

	if err != nil {
		if err == parish.ErrorParishFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *LocationController) GetAllParish(c *gin.Context){

	user, err := u.locationService.GetAllParish()

	if err != nil {
		if err == parish.ErrorParishFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *LocationController) CreateParish(c *gin.Context){
	var location parish.Parish
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.CreateParish(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusCreated, "Ciudad creada satisfactoriamente")
}

func (u *LocationController) UpdateParish(c *gin.Context){
	var location parish.Parish

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.UpdateParish(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Parroquia actualizada satisfactoriamente")
}

func (u *LocationController) DeleteParish(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteParish(id)

	if err != nil {
		if err == city.ErrorCityNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Parroquia eliminada satisfactoriamente")
	return
}

func (u *LocationController) GetStation(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	fire_station, err := u.locationService.GetStation(id)

	if err != nil {
		if err == station.ErrorStationFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fire_station)
	return
}

func (u *LocationController) GetAllStations(c *gin.Context){

	fire_station, err := u.locationService.GetAllStations()

	if err != nil {
		if err == parish.ErrorParishFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fire_station)
	return
}

func (u *LocationController) CreateStation(c *gin.Context){
	var location station.Station
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.CreateStation(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusCreated, "Estación creada satisfactoriamente")
}

func (u *LocationController) UpdateStation(c *gin.Context){
	var location station.Station

	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.locationService.UpdateStation(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Estación actualizada satisfactoriamente")
}

func (u *LocationController) DeleteStation(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.locationService.DeleteParish(id)

	if err != nil {
		if err == station.ErrorStationNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Estación eliminada satisfactoriamente")
	return
}