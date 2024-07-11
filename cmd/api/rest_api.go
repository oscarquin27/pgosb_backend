package api

import (
	auth_handlers "fdms/cmd/api/handlers/auth"
	layout_handlers "fdms/cmd/api/handlers/layouts"
	location_handlers "fdms/cmd/api/handlers/locations"
	roles_handlers "fdms/cmd/api/handlers/roles"
	units_handlers "fdms/cmd/api/handlers/units"
	user_handlers "fdms/cmd/api/handlers/user"
	vehicle_handlers "fdms/cmd/api/handlers/vehicles"
	"fdms/src/infrastructure/config"
	"fdms/src/infrastructure/keycloak"
	"fdms/src/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(db *pgxpool.Pool, auth *keycloak.KeycloakAuthenticationService) {
	router := gin.Default()
	conf := cors.DefaultConfig()

	conf.AllowCredentials = true
	conf.AllowOrigins = []string{"http://localhost:5173", "http://192.168.120.122:5173", "http://192.168.120.110:5173"}

	router.Use(cors.New(conf))

	v1 := router.Group("/api/v1")

	userService := repository.NewUserService(db, auth)
	roleService := repository.NewRoleService(db)

	locationService := repository.NewLocationService(db)
	vehicleService := repository.NewVehicleService(db)
	unityService := repository.NewUnityService(db)
	layoutService := repository.NewLayoutService(db)

	userController := user_handlers.NewUserController(userService)
	roleController := roles_handlers.NewRoleController(roleService)
	locationController := location_handlers.NewLocationController(locationService)
	vehicleController := vehicle_handlers.NewVehicleController(vehicleService)
	unityController := units_handlers.NewUnityController(unityService)
	layoutController := layout_handlers.NewLayoutController(layoutService)

	AuthController := auth_handlers.NewAuthController(auth)

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", AuthController.Login)
		authGroup.PUT("/login", AuthController.RefreshSession)
		authGroup.POST("/logout", AuthController.LogOut)

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
		unity.GET("/:id", unityController.GetUnit)
		unity.GET("/all", unityController.GetAllUnits)
		unity.POST("/create", unityController.CreateUnit)
		unity.PUT("/update", unityController.UpdateUnit)
		unity.DELETE("/:id", unityController.DeleteUnit)
	}

	layout := v1.Group("layout")
	{
		layout.GET("/:entity", layoutController.GetLayout)
	}

	router.Run(":" + config.Configuration.Http.Port)
}
