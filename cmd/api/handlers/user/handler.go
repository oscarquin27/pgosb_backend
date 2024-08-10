package user_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService            services.UserService
	abstractServiceHandler abstract_handler.AbstractHandler[models.User, api_models.UserJson]
}

func NewUserController(userService services.UserService) *UserController {
	abstractHandler := abstract_handler.NewAbstractHandler[models.User, api_models.UserJson](userService)

	return &UserController{
		userService:            userService,
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *UserController) GetUser(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToUserJson, c)
}

func (u *UserController) GetAllUser(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToUserJson, c)

}

func (u *UserController) Create(c *gin.Context) {
	user := api_models.UserJson{}

	var model abstract_handler.AbstactModel[models.User, api_models.UserJson] = &user

	u.abstractServiceHandler.Create(model, api_models.ModelToUserJson, c)
}

func (u *UserController) Update(c *gin.Context) {
	user := api_models.UserJson{}

	var model abstract_handler.AbstactModel[models.User, api_models.UserJson] = &user

	u.abstractServiceHandler.Update(model, api_models.ModelToUserJson, c)
}

func (u *UserController) Delete(c *gin.Context) {

	u.abstractServiceHandler.Delete(c)

}
