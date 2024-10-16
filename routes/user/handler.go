package routes

import (
	entities "fdms/domain/entities/users"
	users "fdms/domain/users"
	"fdms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService users.UserRepository
}

func NewUserController(userService users.UserRepository) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetUser(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

	user, err := u.userService.GetUser(id)

	if err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userDto := u.userService.MapToDto(user)

	c.JSON(http.StatusOK, userDto)

}

func (u *UserController) GetAllUser(c *gin.Context) {

	//time.Sleep(12000 * time.Millisecond)

	user, err := u.userService.GetAll()

	if err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	usersDto := []entities.UserDto{}

	for _, us := range user {
		newUser := u.userService.MapToDto(&us)
		usersDto = append(usersDto, newUser)
	}
	c.JSON(http.StatusOK, usersDto)
}

func (u *UserController) Create(c *gin.Context) {
	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.userService.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Usuario creado satisfactoriamente")
}

func (u *UserController) Update(c *gin.Context) {
	var user entities.User

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

func (u *UserController) Delete(c *gin.Context) {

	id := utils.ParseInt(c.Param("id"))

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

}
