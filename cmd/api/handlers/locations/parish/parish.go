package parish_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type ParishController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Parish, api_models.ParishJson]
}

func NewParishController(municipalityService services.ParishService) *ParishController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Parish, api_models.ParishJson](municipalityService)

	return &ParishController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *ParishController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToParishJson, c)
}

func (u *ParishController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToParishJson, c)
}

func (u *ParishController) Create(c *gin.Context) {
	s := api_models.ParishJson{}

	var model abstract_handler.AbstactModel[models.Parish, api_models.ParishJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToParishJson, c)
}

func (u *ParishController) Update(c *gin.Context) {

	s := api_models.ParishJson{}

	var model abstract_handler.AbstactModel[models.Parish, api_models.ParishJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToParishJson, c)
}

func (u *ParishController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
