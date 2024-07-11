package roles_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (u *RoleController) GetRole(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	role, err := u.roleService.Get(id)

	if err != nil {
		if err == models.ErrorRoleNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	roleResult, _ := api_models.ModelToRoleJson(role)

	c.JSON(http.StatusOK, roleResult)
}

func (u *RoleController) GetAllRoles(c *gin.Context) {

	role, err := u.roleService.GetAll()

	if err != nil {
		if err == models.ErrorRoleNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var rolesResult []api_models.RoleJson

	for _, r := range role {
		newRole, _ := api_models.ModelToRoleJson(&r)
		rolesResult = append(rolesResult, *newRole)
	}

	c.JSON(http.StatusOK, rolesResult)

}

func (u *RoleController) Create(c *gin.Context) {
	var role models.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.roleService.Create(&role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Rol creado satisfactoriamente")
}

func (u *RoleController) Update(c *gin.Context) {
	var role models.Role

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.roleService.Update(&role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Rol actualizado satisfactoriamente")
}

func (u *RoleController) Delete(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.roleService.Delete(id)

	if err != nil {
		if err == models.ErrorRoleNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Rol eliminado satisfactoriamente")
}
