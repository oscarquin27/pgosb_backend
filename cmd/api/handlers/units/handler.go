package units_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UnitController struct {
	service services.UnitService

	abstractServiceHandler abstract_handler.AbstractHandler[models.Unit, api_models.UnitJson]
}

func NewUnityController(unitService services.UnitService) *UnitController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Unit, api_models.UnitJson](unitService)

	return &UnitController{
		service:                unitService,
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *UnitController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToUnitJson, c)
}

func (u *UnitController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToUnitJson, c)
}

func (u *UnitController) Create(c *gin.Context) {
	unit := api_models.UnitJson{}

	var model abstract_handler.AbstactModel[models.Unit, api_models.UnitJson] = &unit

	u.abstractServiceHandler.Create(model, api_models.ModelToUnitJson, c)
}

func (u *UnitController) Update(c *gin.Context) {

	unit := api_models.UnitJson{}

	var model abstract_handler.AbstactModel[models.Unit, api_models.UnitJson] = &unit

	u.abstractServiceHandler.Update(model, api_models.ModelToUnitJson, c)
}

func (u *UnitController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}

func (u *UnitController) GetAllSimple(c *gin.Context) {

	result := u.service.GetAllSimple()

	if !result.IsSuccessful {

		logger.Warn().Err(result.Err.AssociateException()).
			Msg("Problemas ejecutando GetAllSimple")

		if result.Err.Code() == results.NotFoundErr {
			c.JSON(http.StatusOK, result.Value)
		}

		c.JSON(http.StatusInternalServerError, result.Value)
		return
	}

	c.JSON(http.StatusOK, result.Value)
}
