package urbanization_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type UrbanizationController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Urbanization, api_models.UrbanizationJson]
}

func NewUrbanizationController(municipalityService services.UrbanizationService) *UrbanizationController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Urbanization, api_models.UrbanizationJson](municipalityService)

	return &UrbanizationController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *UrbanizationController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToUrbanizationJson, c)
}

func (u *UrbanizationController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToUrbanizationJson, c)
}

func (u *UrbanizationController) Create(c *gin.Context) {
	s := api_models.UrbanizationJson{}

	var model abstract_handler.AbstactModel[models.Urbanization, api_models.UrbanizationJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToUrbanizationJson, c)
}

func (u *UrbanizationController) Update(c *gin.Context) {

	s := api_models.UrbanizationJson{}

	var model abstract_handler.AbstactModel[models.Urbanization, api_models.UrbanizationJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToUrbanizationJson, c)
}

func (u *UrbanizationController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
