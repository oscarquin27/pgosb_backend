package state_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type StateController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.State, api_models.StateJson]
}

func NewStateController(stationService services.StateService) *StateController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.State, api_models.StateJson](stationService)

	return &StateController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *StateController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToStateJson, c)
}

func (u *StateController) GetAll(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToStateJson, c)
}

func (u *StateController) Create(c *gin.Context) {
	s := api_models.StateJson{}

	var model abstract_handler.AbstactModel[models.State, api_models.StateJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToStateJson, c)
}

func (u *StateController) Update(c *gin.Context) {

	s := api_models.StateJson{}

	var model abstract_handler.AbstactModel[models.State, api_models.StateJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToStateJson, c)
}

func (u *StateController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
