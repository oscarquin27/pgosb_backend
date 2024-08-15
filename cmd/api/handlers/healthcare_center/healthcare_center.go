package healthcare_center_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"

	"github.com/gin-gonic/gin"
)

type HealthcareCenterController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.HealthcareCenter, api_models.HealthcareCenterJson]
}

func NewHealthcareCenterController(stationService abstract_handler.AbstractCRUDService[models.HealthcareCenter]) *HealthcareCenterController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.HealthcareCenter, api_models.HealthcareCenterJson](stationService)

	return &HealthcareCenterController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *HealthcareCenterController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToHealthcareCenterJson, c)
}

func (u *HealthcareCenterController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToHealthcareCenterJson, c)
}

func (u *HealthcareCenterController) Create(c *gin.Context) {
	s := api_models.HealthcareCenterJson{}

	var model abstract_handler.AbstactModel[models.HealthcareCenter, api_models.HealthcareCenterJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToHealthcareCenterJson, c)
}

func (u *HealthcareCenterController) Update(c *gin.Context) {

	s := api_models.HealthcareCenterJson{}

	var model abstract_handler.AbstactModel[models.HealthcareCenter, api_models.HealthcareCenterJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToHealthcareCenterJson, c)
}

func (u *HealthcareCenterController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
