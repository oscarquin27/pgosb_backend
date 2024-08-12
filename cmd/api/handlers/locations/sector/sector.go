package sector_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type SectorController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Sector, api_models.SectorJson]
}

func NewSectorController(municipalityService services.SectorService) *SectorController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Sector, api_models.SectorJson](municipalityService)

	return &SectorController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *SectorController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToSectorJson, c)
}

func (u *SectorController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToSectorJson, c)
}

func (u *SectorController) Create(c *gin.Context) {
	s := api_models.SectorJson{}

	var model abstract_handler.AbstactModel[models.Sector, api_models.SectorJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToSectorJson, c)
}

func (u *SectorController) Update(c *gin.Context) {

	s := api_models.SectorJson{}

	var model abstract_handler.AbstactModel[models.Sector, api_models.SectorJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToSectorJson, c)
}

func (u *SectorController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
