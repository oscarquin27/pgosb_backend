package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	layout_service "fdms/domain/layouts"
	location_service "fdms/domain/locations"
	role_service "fdms/domain/roles"
	unity_service "fdms/domain/units"
	user_service "fdms/domain/users"
	vehicle_service "fdms/domain/vehicles"
	config "fdms/infra/config"
	auth_routes "fdms/routes/auth"
	layout_routes "fdms/routes/layouts"
	location_routes "fdms/routes/locations"
	role_routes "fdms/routes/roles"
	unity_routes "fdms/routes/units"
	user_routes "fdms/routes/user"
	vehicle_routes "fdms/routes/vehicles"
)

func Run(db *pgxpool.Pool) {
	router := gin.Default()
	conf := cors.DefaultConfig()

	conf.AllowCredentials = true
	conf.AllowOrigins = []string{"http://localhost:5173", "http://192.168.120.122:5173", "http://192.168.120.110:5173"}

	router.Use(cors.New(conf))

	v1 := router.Group("/api/v1")

	userService := user_service.NewUserService(db)
	roleService := role_service.NewRoleService(db)
	locationService := location_service.NewLocationService(db)
	vehicleService := vehicle_service.NewVehicleService(db)
	unityService := unity_service.NewUnityService(db)
	layoutService := layout_service.NewLayoutService(db)
	userController := user_routes.NewUserController(userService)
	roleController := role_routes.NewRoleController(roleService)
	locationController := location_routes.NewLocationController(locationService)
	vehicleController := vehicle_routes.NewVehicleController(vehicleService)
	unityController := unity_routes.NewUnityController(unityService)
	layoutController := layout_routes.NewLayoutController(layoutService)

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", auth_routes.Login)
		authGroup.PUT("/login", auth_routes.RefreshSession)
		authGroup.POST("/logout", auth_routes.LogOut)

	}

	user := v1.Group("/user")
	{
		user.GET("/:id", userController.GetUser)

		user.GET("/all",
			//auth_routes.PermissionAuthMiddleware(modules.Users, permission.Read, userService, roleService),
			userController.GetAllUser)

		user.POST("/create",
			//auth_routes.PermissionAuthMiddleware(modules.Users, permission.Write, userService, roleService),
			userController.Create)

		user.PUT("/update",
			//auth_routes.PermissionAuthMiddleware(modules.Users, permission.Update, userService, roleService),
			userController.Update)

		user.DELETE("/:id",
			//auth_routes.PermissionAuthMiddleware(modules.Users, permission.Delete, userService, roleService),
			userController.Delete)
	}

	role := v1.Group("/role")
	{
		role.GET("/:id", roleController.GetRole)

		role.GET("/all", roleController.GetAllRoles)
		role.POST("/create", roleController.Create)
		role.PUT("/update", roleController.Update)
		role.DELETE("/:id", roleController.Delete)
	}

	state := v1.Group("/location/state")
	{
		state.GET("/:id", locationController.GetState)
		state.GET("/all", locationController.GetAllStates)
		state.POST("/create", locationController.CreateState)
		state.PUT("/update", locationController.UpdateState)
		state.DELETE("/:id", locationController.DeleteState)
	}

	city := v1.Group("/location/city")
	{
		city.GET("/:id", locationController.GetCity)
		city.GET("/all", locationController.GetAllCity)
		city.POST("/create", locationController.CreateCity)
		city.PUT("/update", locationController.UpdateCity)
		city.DELETE("/:id", locationController.DeleteCity)
	}

	municipality := v1.Group("/location/municipality")
	{
		municipality.GET("/:id", locationController.GetMunicipality)
		municipality.GET("/all", locationController.GetAllMunicipality)
		municipality.POST("/create", locationController.CreateMunicipality)
		municipality.PUT("/update", locationController.UpdateMunicipality)
		municipality.DELETE("/:id", locationController.DeleteMunicipality)
	}

	parish := v1.Group("/location/parish")
	{
		parish.GET("/:id", locationController.GetParish)
		parish.GET("/all", locationController.GetAllParish)
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

	vehicle := v1.Group("/vehicles")
	{
		vehicle.GET("/:id", vehicleController.GetVehicle)
		vehicle.GET("/all", vehicleController.GetAllVehicle)
		vehicle.POST("/create", vehicleController.CreateVehicle)
		vehicle.PUT("/update", vehicleController.UpdateVehicle)
		vehicle.DELETE("/:id", vehicleController.DeleteVehicle)
	}

	unity := v1.Group("unit")
	{
		unity.GET("/:id", unityController.GetUnity)
		unity.GET("/all", unityController.GetAllUnities)
		unity.POST("/create", unityController.CreateUnity)
		unity.PUT("/update", unityController.UpdateUnity)
		unity.DELETE("/:id", unityController.DeleteUnity)
	}

	layout := v1.Group("layout")
	{
		layout.GET("/:entity", layoutController.GetLayout)
	}

	router.Run(":" + config.Configuration.Http.Port)
}
