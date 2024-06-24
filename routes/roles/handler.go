package routes

import (
	entities "fdms/domain/entities/roles"
	roles "fdms/domain/roles"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService roles.RoleRepository
}

func NewRoleController(roleService roles.RoleRepository) *RoleController {
	return &RoleController{
		roleService : roleService,
	}
}

func (u *RoleController) GetRole(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.roleService.GetRole(id)

	if err != nil {
		if err == entities.ErrorRoleNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *RoleController) GetAllRoles(c *gin.Context){

	user, err := u.roleService.GetAll()

	if err != nil {
		if err == entities.ErrorRoleNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *RoleController) Create(c *gin.Context){
	var role entities.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.roleService.Create(&role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Rol creado satisfactoriamente")
}


func (u *RoleController) Update(c *gin.Context){
	var role entities.Role

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.roleService.Update(&role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Rol actualizado satisfactoriamente")
}

func (u *RoleController) Delete(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.roleService.Delete(id)

	if err != nil {
		if err == entities.ErrorRoleNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Rol eliminado satisfactoriamente")
	return
}
