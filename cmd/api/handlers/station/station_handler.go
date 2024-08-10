package station_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"

	"github.com/gin-gonic/gin"
)

type StationController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Station, api_models.StationJson]
}

func NewStationController(stationService abstract_handler.AbstractCRUDService[models.Station]) *StationController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Station, api_models.StationJson](stationService)

	return &StationController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *StationController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToStationJson, c)
}

func (u *StationController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToStationJson, c)
}

func (u *StationController) Create(c *gin.Context) {
	s := api_models.StationJson{}

	var model abstract_handler.AbstactModel[models.Station, api_models.StationJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToStationJson, c)
}

func (u *StationController) Update(c *gin.Context) {

	s := api_models.StationJson{}

	var model abstract_handler.AbstactModel[models.Station, api_models.StationJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToStationJson, c)
}

func (u *StationController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
