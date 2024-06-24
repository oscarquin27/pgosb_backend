package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	location_service "fdms/domain/locations"
	role_service "fdms/domain/roles"
	unity_service "fdms/domain/unities"
	user_service "fdms/domain/users"
	vehicle_service "fdms/domain/vehicles"
	location_routes "fdms/routes/locations"
	role_routes "fdms/routes/roles"
	unity_routes "fdms/routes/unities"
	user_routes "fdms/routes/user"
	vehicle_routes "fdms/routes/vehicles"
)

var router *gin.Engine

func Run(db *pgxpool.Pool) {
  router := gin.Default()
  v1 := router.Group("/api/v1")

  userService := user_service.NewUserService(db)
  roleService := role_service.NewRoleService(db)
  locationService := location_service.NewLocationService(db)
  vehicleService := vehicle_service.NewVehicleService(db)
  unityService := unity_service.NewUnityService(db)
  userController := user_routes.NewUserController(userService)
  roleController := role_routes.NewRoleController(roleService)
  locationController := location_routes.NewLocationController(locationService)
  vehicleController := vehicle_routes.NewVehicleController(vehicleService)
  unityController := unity_routes.NewUnityController(unityService)
  
  user := v1.Group("/user")
  {
    user.GET("/:id", userController.GetUser)
    user.GET("all", userController.GetAllUser)
    user.POST("/create", userController.Create)
    user.PUT("/update", userController.Update)
    user.DELETE("/:id", userController.Delete)
  }

  role := v1.Group("/role")
  {
    role.GET("/:id", roleController.GetRole) 
    role.GET("all", roleController.GetAllRoles) 	
    role.POST("/create", roleController.Create) 
    role.PUT("/update", roleController.Update)  
    role.DELETE("/:id", roleController.Delete)  
  }

  state := v1.Group("/location/state")
  {
    state.GET("/:id", locationController.GetState)
    state.GET("all", locationController.GetAllStates)
    state.POST("/create", locationController.CreateState)
    state.PUT("/update", locationController.UpdateState)
    state.DELETE("/:id", locationController.DeleteState)
  }

  city := v1.Group("/location/city")
  {
    city.GET("/:id", locationController.GetCity)
    city.GET("all", locationController.GetAllCity)
    city.POST("/create", locationController.CreateCity)
    city.PUT("/update", locationController.UpdateCity)
    city.DELETE("/:id", locationController.DeleteCity)
  }

  municipality := v1.Group("/location/municipality")
  {
    municipality.GET("/:id", locationController.GetMunicipality)
    municipality.GET("all", locationController.GetAllMunicipality)
    municipality.POST("/create", locationController.CreateMunicipality)
    municipality.PUT("/update", locationController.UpdateMunicipality)
    municipality.DELETE("/:id", locationController.DeleteMunicipality)
  }

  parish := v1.Group("/location/parish")
  {
    parish.GET("/:id", locationController.GetParish)
    parish.GET("all", locationController.GetAllParish)
    parish.POST("/create", locationController.CreateParish)
    parish.PUT("/update", locationController.UpdateParish)
    parish.DELETE("/:id", locationController.DeleteParish)
  }

  station := v1.Group("/location/station")
  {
    station.GET("/:id", locationController.GetStation)
    station.GET("all", locationController.GetAllStations)
    station.POST("/create", locationController.CreateStation)
    station.PUT("/update", locationController.UpdateStation)
    station.DELETE("/:id", locationController.DeleteStation)   
  }
  
  vehicle := v1.Group("vehicles")
  {
    vehicle.GET("/:id", vehicleController.GetVehicle)
    vehicle.GET("all", vehicleController.GetAllVehicle)
    vehicle.POST("/create", vehicleController.CreateVehicle)
    vehicle.PUT("/update", vehicleController.UpdateVehicle)
    vehicle.DELETE("/:id", vehicleController.DeleteVehicle)  
  }

  unity := v1.Group("unity")
  {
    unity.GET("/:id", unityController.GetUnity)
    unity.GET("all", unityController.GetAllUnities)
    unity.POST("/create", unityController.CreateUnity)
    unity.PUT("/update", unityController.UpdateUnity)
    unity.DELETE("/:id", unityController.DeleteUnity)  
  }

  router.Run(":5000")
}
