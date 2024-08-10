package roles_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService            services.RoleService
	abstractServiceHandler abstract_handler.AbstractHandler[models.Role, api_models.RoleJson]
}

func NewRoleController(roleService services.RoleService) *RoleController {
	abstractHandler := abstract_handler.NewAbstractHandler[models.Role, api_models.RoleJson](roleService)

	return &RoleController{
		roleService:            roleService,
		abstractServiceHandler: *abstractHandler,
	}
}

func (u *RoleController) GetRole(c *gin.Context) {
	u.abstractServiceHandler.Get(api_models.ModelToRoleJson, c)
}

func (u *RoleController) GetAllRoles(c *gin.Context) {

	u.abstractServiceHandler.GetAll(api_models.ModelToRoleJson, c)
}
func (u *RoleController) Create(c *gin.Context) {
	role := api_models.RoleJson{}

	var model abstract_handler.AbstactModel[models.Role, api_models.RoleJson] = &role

	u.abstractServiceHandler.Create(model, api_models.ModelToRoleJson, c)
}

func (u *RoleController) Update(c *gin.Context) {
	role := api_models.RoleJson{}

	var model abstract_handler.AbstactModel[models.Role, api_models.RoleJson] = &role

	u.abstractServiceHandler.Update(model, api_models.ModelToRoleJson, c)
}

func (u *RoleController) Delete(c *gin.Context) {
	u.abstractServiceHandler.Delete(c)
}
