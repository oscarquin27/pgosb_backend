package api

import (
	auth_handlers "fdms/cmd/api/handlers/auth"
	health_check "fdms/cmd/api/handlers/health_cheack"
	healthcare_center_handler "fdms/cmd/api/handlers/healthcare_center"
	layout_handlers "fdms/cmd/api/handlers/layouts"
	municipality_handler "fdms/cmd/api/handlers/locations/municipality"
	parish_handler "fdms/cmd/api/handlers/locations/parish"
	sector_handler "fdms/cmd/api/handlers/locations/sector"
	state_handler "fdms/cmd/api/handlers/locations/states"
	urbanization_handler "fdms/cmd/api/handlers/locations/urbanization"
	mission_handlers "fdms/cmd/api/handlers/mission"
	antares_handlers "fdms/cmd/api/handlers/mission_antares"
	mission_authority_handler "fdms/cmd/api/handlers/mission_authority"
	mission_firefighter_handler "fdms/cmd/api/handlers/mission_firefghter"
	mission_infra_handlers "fdms/cmd/api/handlers/mission_infrastructure"
	mission_location_handler "fdms/cmd/api/handlers/mission_location"
	mission_person_handlers "fdms/cmd/api/handlers/mission_person"
	mission_service_handlers "fdms/cmd/api/handlers/mission_services"
	mission_vehicle_handlers "fdms/cmd/api/handlers/mission_vehicles"
	operative_regions_handlers "fdms/cmd/api/handlers/operative_regions"
	roles_handlers "fdms/cmd/api/handlers/roles"
	station_handler "fdms/cmd/api/handlers/station"

	authority_handler "fdms/cmd/api/handlers/authority"

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
	healthcareCenterService := repository.NewHealthcareCenterService(db)
	stateService := repository.NewStateService(db)
	operativeRegionService := repository.NewOPerativeRegionsService(db)

	authorityService := repository.NewAuthorityService(db)

	municipalityService := repository.NewMunicipalityService(db)
	parishSevice := repository.NewParishService(db)
	sectorService := repository.NewSectorService(db)
	urbanizationService := repository.NewUrbanizationService(db)

	vehicleService := repository.NewVehicleService(db)

	layoutService := repository.NewLayoutService(db)

	missionService := repository.NewMissionService(db)
	missionServiceService := repository.NewMissionServiceService(db)
	missionVehicleService := repository.NewMissionVehicleService(db)
	missionPersonService := repository.NewMissionPersonService(db)
	missionInfraService := repository.NewMissionInfrastructureService(db)
	missionAntaresService := repository.NewAntaresService(db)

	missionFireFighterService := repository.NewMissionFirefighterService(db)

	missionLocationService := repository.NewMissionLocationService(db)

	missionAuthorityService := repository.NewMissionAuthorityService(db)
	missionAuthorityVehicleService := repository.NewMissionAuthorityVehicleService(db)

	missionAuthorityPersonService := repository.NewMissionAuthorityPersonService(db)

	missionAuthorityServiceService := repository.NewMissionAuthorityServiceService(db)

	authorityController := authority_handler.NewAuthorityController(authorityService)

	missionController := mission_handlers.NewMissionController(missionService)

	missionAuthorityVehicleController := mission_authority_handler.NewMissionAuthorityVehicleController(missionAuthorityVehicleService, missionAuthorityVehicleService)
	missionAuthorityPersonController := mission_authority_handler.NewMissionAuthorityPersonController(missionAuthorityPersonService, missionAuthorityPersonService)
	missionAuthorityServiceController := mission_authority_handler.NewMissionAuthorityServiceController(missionAuthorityServiceService, missionAuthorityServiceService)

	missionAuthorityController := mission_authority_handler.NewMissionAuthorityController(missionAuthorityService, missionAuthorityService)

	missionServiceController := mission_service_handlers.NewServiceServiceController(missionServiceService)
	missionVehicleController := mission_vehicle_handlers.NewMissionVehicleController(missionVehicleService)
	missionPersonController := mission_person_handlers.NewMissionPersonController(missionPersonService)
	missionInfraController := mission_infra_handlers.NewMissionController(missionInfraService)

	missionLocationController := mission_location_handler.NewMissionLocationController(missionLocationService, missionLocationService)
	missionFireFighterController := mission_firefighter_handler.NewMissionFireFigtherController(missionFireFighterService)

	userController := user_handlers.NewUserController(userService)
	roleController := roles_handlers.NewRoleController(roleService)
	unityController := units_handlers.NewUnityController(unityService)
	stationController := station_handler.NewStationController(stationService)
	healthCareCenterController := healthcare_center_handler.NewHealthcareCenterController(healthcareCenterService)
	operativeRegionController := operative_regions_handlers.NewOperativeRegionController(operativeRegionService)
	stateController := state_handler.NewStateController(stateService)
	municpalityController := municipality_handler.NewMunicipalityController(municipalityService)
	parishController := parish_handler.NewParishController(parishSevice)
	sectorController := sector_handler.NewSectorController(sectorService)
	urbanizationController := urbanization_handler.NewUrbanizationController(urbanizationService)

	vehicleController := vehicle_handlers.NewVehicleController(vehicleService)

	missionAntaresController := antares_handlers.NewAntaresController(missionAntaresService)
	layoutController := layout_handlers.NewLayoutController(layoutService)

	healthCheckController := health_check.NewHealthCheckService(db, auth)

	AuthController := auth_handlers.NewAuthController(auth)

	conf.AllowCredentials = true

	conf.AllowOrigins = []string{"https://gres.local.net:8083", "http://localhost:5173", "http://192.168.120.136:5173"}

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: logger.Log(),
	}))

	router.Use(ZerologMiddleware())
	router.Use(cors.New(conf))

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", healthCheckController.HealthCheckHandler)

	v1 := router.Group("/api/v1")

	authGroup := v1.Group("/auth")
	{
		authGroup.GET("/login/test", AuthController.LoginTest)
		authGroup.POST("/login", AuthController.Login)
		authGroup.PUT("/login", AuthController.RefreshSession)
		authGroup.POST("/logout", AuthController.LogOut)
		authGroup.POST("/change-password", AuthController.ChangePassword)

	}

	authority := v1.Group("/authority")
	{
		authority.GET("/:id", authorityController.Get)
		authority.GET("/all", authorityController.GetAll)
		authority.POST("/create", authorityController.Create)
		authority.PUT("/update", authorityController.Update)
		authority.DELETE("/:id", authorityController.Delete)
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

	municipality := v1.Group("/location/municipality")
	{
		municipality.GET("/:id", municpalityController.Get)
		municipality.GET("/all", municpalityController.GetAll)
		municipality.POST("/create", municpalityController.Create)
		municipality.PUT("/update", municpalityController.Update)
		municipality.DELETE("/:id", municpalityController.Delete)
	}

	parish := v1.Group("/location/parish")
	{
		parish.GET("/:id", parishController.Get)
		parish.GET("/all", parishController.GetAll)
		parish.POST("/create", parishController.Create)
		parish.PUT("/update", parishController.Update)
		parish.DELETE("/:id", parishController.Delete)
	}
	sector := v1.Group("/location/sector")
	{
		sector.GET("/:id", sectorController.Get)
		sector.GET("/all", sectorController.GetAll)
		sector.POST("/create", sectorController.Create)
		sector.PUT("/update", sectorController.Update)
		sector.DELETE("/:id", sectorController.Delete)
	}

	urbanization := v1.Group("/location/urbanization")
	{
		urbanization.GET("/:id", urbanizationController.Get)
		urbanization.GET("/all", urbanizationController.GetAll)
		urbanization.POST("/create", urbanizationController.Create)
		urbanization.PUT("/update", urbanizationController.Update)
		urbanization.DELETE("/:id", urbanizationController.Delete)
	}

	station := v1.Group("/station")
	{
		station.GET("/:id", stationController.Get)
		station.GET("all", stationController.GetAll)
		station.POST("/create", stationController.Create)
		station.PUT("/update", stationController.Update)
		station.DELETE("/:id", stationController.Delete)
	}

	centers := v1.Group("center")
	{
		centers.GET("/:id", healthCareCenterController.Get)
		centers.GET("/all", healthCareCenterController.GetAll)
		centers.POST("/create", healthCareCenterController.Create)
		centers.PUT("/update", healthCareCenterController.Update)
		centers.DELETE("/:id", healthCareCenterController.Delete)
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
		serviceMission.GET("/summary", missionServiceController.GetAllSummary)
		serviceMission.GET("/relevant/:id", missionServiceController.GetRelevantServices)
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

	firefightersMission := v1.Group("mission/firefighter")
	{
		firefightersMission.GET("/:id", missionFireFighterController.Get)
		firefightersMission.POST("/create", missionFireFighterController.Create)
		firefightersMission.PUT("/update", missionFireFighterController.Update)
		firefightersMission.DELETE("/delete/:id", missionFireFighterController.Delete)

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

	missionAuthority := v1.Group("mission/authority")
	{
		missionAuthority.GET("/:id", missionAuthorityController.Get)
		missionAuthority.GET("/group/:id", missionAuthorityController.GetByMissionId)
		missionAuthority.GET("/all", missionAuthorityController.GetAll)
		missionAuthority.POST("/create", missionAuthorityController.Create)
		missionAuthority.PUT("/update", missionAuthorityController.Update)
		missionAuthority.DELETE("/delete/:id", missionAuthorityController.Delete)
		missionAuthority.GET("/summary/:id", missionAuthorityController.GetSummaryByMissionId)
	}

	missionAuthorityVehicle := v1.Group("mission/authority/vehicle")
	{
		missionAuthorityVehicle.GET("/:id", missionAuthorityVehicleController.Get)
		missionAuthorityVehicle.GET("/group/:id", missionAuthorityVehicleController.GetByAuthorityId)
		missionAuthorityVehicle.GET("/all", missionAuthorityVehicleController.GetAll)
		missionAuthorityVehicle.POST("/create", missionAuthorityVehicleController.Create)
		missionAuthorityVehicle.PUT("/update", missionAuthorityVehicleController.Update)
		missionAuthorityVehicle.DELETE("/delete/:id", missionAuthorityVehicleController.Delete)
	}

	missionAuthorityPerson := v1.Group("mission/authority/person")
	{
		missionAuthorityPerson.GET("/:id", missionAuthorityPersonController.Get)
		missionAuthorityPerson.GET("/group/:id", missionAuthorityPersonController.GetByAuthorityId)
		missionAuthorityPerson.GET("/all", missionAuthorityPersonController.GetAll)
		missionAuthorityPerson.POST("/create", missionAuthorityPersonController.Create)
		missionAuthorityPerson.PUT("/update", missionAuthorityPersonController.Update)
		missionAuthorityPerson.DELETE("/delete/:id", missionAuthorityPersonController.Delete)
	}

	missionAuthorityServiceRelated := v1.Group("mission/authority/service")
	{
		missionAuthorityServiceRelated.GET("/:id", missionAuthorityServiceController.Get)
		missionAuthorityServiceRelated.GET("/group/:id", missionAuthorityServiceController.GetByServiceId)
		missionAuthorityServiceRelated.GET("/all", missionAuthorityServiceController.GetAll)
		missionAuthorityServiceRelated.POST("/create", missionAuthorityServiceController.Create)
		missionAuthorityServiceRelated.PUT("/update", missionAuthorityServiceController.Update)
		missionAuthorityServiceRelated.DELETE("/delete/:id", missionAuthorityServiceController.Delete)
	}

	locationMission := v1.Group("mission/location")

	{
		locationMission.GET("/:id", missionLocationController.Get)
		locationMission.GET("/group/:id", missionLocationController.GetLocationsByServiceId)
		locationMission.GET("/all", missionLocationController.GetAll)
		locationMission.POST("/create", missionLocationController.Create)
		locationMission.PUT("/update", missionLocationController.Update)
		locationMission.DELETE("/delete/:id", missionLocationController.Delete)
	}

	operativeRegion := v1.Group("operative/region")
	{
		operativeRegion.GET("/:id", operativeRegionController.Get)
		operativeRegion.GET("/all", operativeRegionController.GetAll)
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
