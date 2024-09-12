package operative_region_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"

	"github.com/gin-gonic/gin"
)

type OperativeRegionController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.OperativeRegions, api_models.OperativeRegionsJson]
}

func NewOperativeRegionController(operativeRegionService abstract_handler.AbstractCRUDService[models.OperativeRegions]) *OperativeRegionController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.OperativeRegions, api_models.OperativeRegionsJson](operativeRegionService)

	return &OperativeRegionController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *OperativeRegionController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToOperativeRegionsJson, c)
}

func (u *OperativeRegionController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToOperativeRegionsJson, c)
}
