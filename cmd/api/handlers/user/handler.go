package user_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetUser(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := u.userService.Get(id)

	if err != nil {
		if err == models.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userDto := api_models.ModelToUserJson(user)

	c.JSON(http.StatusOK, userDto)

}

func (u *UserController) GetAllUser(c *gin.Context) {

	//time.Sleep(12000 * time.Millisecond)

	user, err := u.userService.GetAll()

	if err != nil {
		if err == models.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	usersDto := []api_models.UserJson{}

	for _, us := range user {
		newUser := api_models.ModelToUserJson(&us)
		usersDto = append(usersDto, *newUser)
	}
	c.JSON(http.StatusOK, usersDto)
}

func (u *UserController) Create(c *gin.Context) {
	var user models.User

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
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.userService.Update(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Usuario actualizado satisfactoriamente")
}

func (u *UserController) Delete(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.userService.Delete(id)

	if err != nil {
		if err == models.ErrorUserNotDeleted {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Usuario eliminado satisfactoriamente")

}
