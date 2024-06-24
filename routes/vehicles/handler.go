package routes

import (
	entities "fdms/domain/entities/vehicles"
	vehicle "fdms/domain/vehicles"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	vehicleService vehicle.VehicleRepository
}

func NewVehicleController(vehicleService vehicle.VehicleRepository) *VehicleController {
	return &VehicleController{
		vehicleService : vehicleService,
	}
}

func (u *VehicleController) GetVehicle(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	vehicle, err := u.vehicleService.GetVehicle(id)

	if err != nil {
		if err == entities.ErrorVehicleNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
	return
}

func (u *VehicleController) GetAllVehicle(c *gin.Context){

	vehicle, err := u.vehicleService.GetAll()

	if err != nil {
		if err == entities.ErrorVehicleNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vehicle)
	return
}

func (u *VehicleController) CreateVehicle(c *gin.Context){
	var user entities.Vehicle
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.vehicleService.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Vehículo creado satisfactoriamente")
}


func (u *VehicleController) UpdateVehicle(c *gin.Context){
	var user entities.Vehicle

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.vehicleService.Update(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Vehículo actualizado satisfactoriamente")
}

func (u *VehicleController) DeleteVehicle(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.vehicleService.Delete(id)

	if err != nil {
		if err == entities.ErrorVehicleNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Vehículo eliminado satisfactoriamente")
	return
}
