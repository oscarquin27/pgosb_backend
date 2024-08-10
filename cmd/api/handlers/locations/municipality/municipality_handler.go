package municipality_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type MunicipalityController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Municipality, api_models.MunicipalityJson]
}

func NewMunicipalityController(municipalityService services.MunicipalityService) *MunicipalityController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Municipality, api_models.MunicipalityJson](municipalityService)

	return &MunicipalityController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *MunicipalityController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToMunicipalityJson, c)
}

func (u *MunicipalityController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToMunicipalityJson, c)
}

func (u *MunicipalityController) Create(c *gin.Context) {
	s := api_models.MunicipalityJson{}

	var model abstract_handler.AbstactModel[models.Municipality, api_models.MunicipalityJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToMunicipalityJson, c)
}

func (u *MunicipalityController) Update(c *gin.Context) {

	s := api_models.MunicipalityJson{}

	var model abstract_handler.AbstactModel[models.Municipality, api_models.MunicipalityJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToMunicipalityJson, c)
}

func (u *MunicipalityController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
