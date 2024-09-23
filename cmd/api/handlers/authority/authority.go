package authority_handler

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"

	"github.com/gin-gonic/gin"
)

type AuthorityController struct {
	abstractServiceHandler abstract_handler.AbstractHandler[models.Authority, api_models.AuthorityJson]
}

func NewAuthorityController(service abstract_handler.AbstractCRUDService[models.Authority]) *AuthorityController {

	abstractHandler := abstract_handler.NewAbstractHandler[models.Authority, api_models.AuthorityJson](service)

	return &AuthorityController{
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *AuthorityController) Get(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToAuthorityJson, c)
}

func (u *AuthorityController) GetAll(c *gin.Context) {
	u.abstractServiceHandler.GetAll(api_models.ModelToAuthorityJson, c)
}

func (u *AuthorityController) Create(c *gin.Context) {
	s := api_models.AuthorityJson{}

	var model abstract_handler.AbstactModel[models.Authority, api_models.AuthorityJson] = &s

	u.abstractServiceHandler.Create(model, api_models.ModelToAuthorityJson, c)

}

func (u *AuthorityController) Update(c *gin.Context) {

	s := api_models.AuthorityJson{}

	var model abstract_handler.AbstactModel[models.Authority, api_models.AuthorityJson] = &s

	u.abstractServiceHandler.Update(model, api_models.ModelToAuthorityJson, c)
}

func (u *AuthorityController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)
}
