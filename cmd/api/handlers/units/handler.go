package units_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"

	"github.com/gin-gonic/gin"
)

type UnitController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Unit, api_models.UnitJson]
}

func NewUnityController(unitService abstract_handler.AbstractCRUDService[models.Unit]) *UnitController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Unit, api_models.UnitJson](unitService)

	return &UnitController{
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
