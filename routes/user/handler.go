package routes

import (
	entities "fdms/domain/entities/users"
	users "fdms/domain/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService users.UserRepository
}

func NewUserController(userService users.UserRepository) *UserController {
	return &UserController{
		userService : userService,
	}
}

func (u *UserController) GetUser(c *gin.Context){

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.userService.GetUser(id)

	if err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *UserController) GetAllUser(c *gin.Context){

	user, err := u.userService.GetAll()

	if err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u *UserController) Create(c *gin.Context){
	var user entities.UserCreateDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.userService.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Usuario creado satisfactoriamente")
}


func (u *UserController) Update(c *gin.Context){
	var user entities.UserUpdateDto

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.userService.Update(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, "Usuario actualizado satisfactoriamente")
}

func (u *UserController) Delete(c *gin.Context){

	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.userService.Delete(id)

	if err != nil {
		if err == entities.ErrorUserNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Usuario eliminado satisfactoriamente")
	return
}
