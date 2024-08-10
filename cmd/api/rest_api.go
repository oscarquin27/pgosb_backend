package api

import (
	auth_handlers "fdms/cmd/api/handlers/auth"
	center_handlers "fdms/cmd/api/handlers/centers"
	layout_handlers "fdms/cmd/api/handlers/layouts"
	location_handlers "fdms/cmd/api/handlers/locations"
	state_handler "fdms/cmd/api/handlers/locations/states"
	mission_handlers "fdms/cmd/api/handlers/mission"
	antares_handlers "fdms/cmd/api/handlers/mission_antares"
	mission_infra_handlers "fdms/cmd/api/handlers/mission_infrastructure"
	mission_person_handlers "fdms/cmd/api/handlers/mission_person"
	mission_service_handlers "fdms/cmd/api/handlers/mission_services"
	mission_vehicle_handlers "fdms/cmd/api/handlers/mission_vehicles"
	roles_handlers "fdms/cmd/api/handlers/roles"
	station_handler "fdms/cmd/api/handlers/station"
	units_handlers "fdms/cmd/api/handlers/units"
	user_handlers "fdms/cmd/api/handlers/user"
	vehicle_handlers "fdms/cmd/api/handlers/vehicles"
	"fdms/src/infrastructure/config"
	"fdms/src/infrastructure/keycloak"
	logger "fdms/src/infrastructure/log"
	"fdms/src/repository"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ZerologMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()

		// Call the next handler
		c.Next()

		// Log the request
		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Dur("latency", time.Since(start)).
			Int("status", c.Writer.Status()).
			Msg("")
	}
}

func Run(db *pgxpool.Pool, auth *keycloak.KeycloakAuthenticationService) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	conf := cors.DefaultConfig()

	userService := repository.NewUserService(db, auth)
	roleService := repository.NewRoleService(db)
	unityService := repository.NewUnityService(db)
	stationService := repository.NewStationService(db)

	stateService := repository.NewStateService(db)

	locationService := repository.NewLocationService(db)
	vehicleService := repository.NewVehicleService(db)

	layoutService := repository.NewLayoutService(db)

	missionService := repository.NewMissionService(db)
	missionServiceService := repository.NewMissionServiceService(db)
	missionVehicleService := repository.NewMissionVehicleService(db)
	missionPersonService := repository.NewMissionPersonService(db)
	missionInfraService := repository.NewMissionInfrastructureService(db)
	missionAntaresService := repository.NewAntaresService(db)

	centerService := repository.NewCenterService(db)

	missionController := mission_handlers.NewMissionController(missionService)
	missionServiceController := mission_service_handlers.NewServiceServiceController(missionServiceService)
	missionVehicleController := mission_vehicle_handlers.NewMissionVehicleController(missionVehicleService)
	missionPersonController := mission_person_handlers.NewMissionPersonController(missionPersonService)
	missionInfraController := mission_infra_handlers.NewMissionController(missionInfraService)

	userController := user_handlers.NewUserController(userService)
	roleController := roles_handlers.NewRoleController(roleService)
	unityController := units_handlers.NewUnityController(unityService)
	stationController := station_handler.NewStationController(stationService)
	stateController := state_handler.NewStateController(stateService)

	locationController := location_handlers.NewLocationController(locationService)

	vehicleController := vehicle_handlers.NewVehicleController(vehicleService)

	centerController := center_handlers.NewCenterController(centerService)
	missionAntaresController := antares_handlers.NewAntaresController(missionAntaresService)
	layoutController := layout_handlers.NewLayoutController(layoutService)

	AuthController := auth_handlers.NewAuthController(auth)

	conf.AllowCredentials = true
	conf.AllowOrigins = []string{"http://localhost:5173",
		"http://192.168.120.122:5173", "http://192.168.0.164:5173", "http://192.168.120.110:5173", "http://192.168.1.12:5100",
		"http://172.30.100.9:8082", "http://192.168.1.12:5173", "http://192.168.1.7:5173", "http://pruebas.gres.local.net:5173"}

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: logger.Log(),
	}))

	router.Use(ZerologMiddleware())
	router.Use(cors.New(conf))

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := router.Group("/api/v1")

	authGroup := v1.Group("/auth")
	{
		authGroup.GET("/login/test", AuthController.LoginTest)
		authGroup.POST("/login", AuthController.Login)
		authGroup.PUT("/login", AuthController.RefreshSession)
		authGroup.POST("/logout", AuthController.LogOut)

	}
	user := v1.Group("/user")
	{
		user.GET("/:id", userController.GetUser)
		user.GET("/all", userController.GetAllUser)
		user.GET("/all/simple", userController.GetAllSimple)

		user.POST("/create", userController.Create)
		user.PUT("/update", userController.Update)
		user.DELETE("/:id", userController.Delete)
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
		state.GET("/:id", stateController.Get)
		state.GET("/all", stateController.GetAll)
		state.POST("/create", stateController.Create)
		state.PUT("/update", stateController.Update)
		state.DELETE("/:id", stateController.Delete)
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

	station := v1.Group("/station")
	{
		station.GET("/:id", stationController.Get)
		station.GET("all", stationController.GetAll)
		station.POST("/create", stationController.Create)
		station.PUT("/update", stationController.Update)
		station.DELETE("/:id", stationController.Delete)
	}

	vehicle := v1.Group("/vehicles")
	{
		vehicle.GET("/:id", vehicleController.GetVehicle)
		vehicle.GET("/all", vehicleController.GetAllVehicle)
		vehicle.POST("/create", vehicleController.CreateVehicle)
		vehicle.PUT("/update", vehicleController.UpdateVehicle)
		vehicle.DELETE("/:id", vehicleController.DeleteVehicle)
		vehicle.GET("/types", vehicleController.GetVehicleType)
		vehicle.POST("/types", vehicleController.GetVehicleModel)
	}

	unity := v1.Group("unit")
	{
		unity.GET("/:id", unityController.Get)
		unity.GET("/all", unityController.GetAll)
		unity.GET("/all/simple", unityController.GetAllSimple)

		unity.POST("/create", unityController.Create)
		unity.PUT("/update", unityController.Update)
		unity.DELETE("/:id", unityController.Delete)
	}

	layout := v1.Group("layout")
	{
		layout.GET("/:entity", layoutController.GetLayout)
	}

	mission := v1.Group("mission")
	{
		mission.GET("/:id", missionController.GetMission)
		mission.GET("/all", missionController.GetAllMissions)
		mission.POST("/create", missionController.Create)
		mission.PUT("/update", missionController.Update)
		mission.DELETE("/:id", missionController.Delete)
	}

	antaresMission := v1.Group("mission/antares")
	{
		antaresMission.GET("/all", missionAntaresController.GetAll)
	}

	serviceMission := v1.Group("mission/service")
	{
		serviceMission.GET("/all", missionServiceController.GetAll)
		serviceMission.GET("/:id", missionServiceController.Get)
		serviceMission.GET("/group/:id", missionServiceController.GetByMissionId)
		serviceMission.POST("/create", missionServiceController.Create)
		serviceMission.PUT("/update", missionServiceController.Update)
		serviceMission.DELETE("/delete/:id", missionServiceController.Delete)
		serviceMission.GET("/unit/:id", missionServiceController.GetUnits)
		serviceMission.GET("/user/:id", missionServiceController.GetUsers)
	}

	vehicleMission := v1.Group("mission/vehicle")
	{
		vehicleMission.GET("/:id", missionVehicleController.GetVehicle)
		vehicleMission.GET("/group/:id", missionVehicleController.GetByServiceId)
		vehicleMission.GET("/all", missionVehicleController.GetAll)
		vehicleMission.POST("/create", missionVehicleController.Create)
		vehicleMission.PUT("/update", missionVehicleController.Update)
		vehicleMission.DELETE("/delete/:id", missionVehicleController.Delete)
	}

	infraMission := v1.Group("mission/infrastructure")
	{
		infraMission.GET("/:id", missionInfraController.GetInfrastructure)
		infraMission.GET("/group/:id", missionInfraController.GetByServiceId)
		infraMission.GET("/all", missionInfraController.GetAll)
		infraMission.POST("/create", missionInfraController.Create)
		infraMission.PUT("/update", missionInfraController.Update)
		infraMission.DELETE("/delete/:id", missionInfraController.Delete)
	}

	personMission := v1.Group("mission/person")
	{
		personMission.GET("/:id", missionPersonController.Get)
		personMission.GET("/group/:id", missionPersonController.GetByServiceId)
		personMission.GET("/all", missionPersonController.GetAll)
		personMission.POST("/create", missionPersonController.Create)
		personMission.PUT("/update", missionPersonController.Update)
		personMission.DELETE("/delete/:id", missionPersonController.Delete)
	}

	centers := v1.Group("center")
	{
		centers.GET("/:id", centerController.GetCenter)
		centers.GET("/all", centerController.GetAllCenters)
		centers.POST("/create", centerController.Create)
		centers.PUT("/update", centerController.Update)
		centers.DELETE("/:id", centerController.Delete)
	}

	if config.Get().Http.EnabledSsl {
		if err := router.RunTLS(fmt.Sprintf("0.0.0.0:%d",
			config.Get().Http.Port),
			config.Get().Http.SslCert,
			config.Get().Http.SslKey); err != nil {

			logger.Fatal().Err(err).Msg("Failed to start REST API server")
		}
	} else if err := router.Run(fmt.Sprintf("0.0.0.0:%d",
		config.Get().Http.Port)); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start REST API server")
	}

}
