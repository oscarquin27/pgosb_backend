package abstract_handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AbstractCRUDService[T any] interface {
	Get(id int64) (*T, error)
	GetAll() ([]T, error)
	Create(value *T) error
	Update(value *T) error
	Delete(id int64) error
}

type AbstractHandler[T any, F any] struct {
	service AbstractCRUDService[T]
}

// type FromModelFunctionType[T any, F any] func(*T) *F

func NewAbstractHandler[T any, F any](s AbstractCRUDService[T]) *AbstractHandler[T, F] {

	return &AbstractHandler[T, F]{
		service: s,
	}
}

func (u *AbstractHandler[T, F]) Get(FromModel func(*T) *F,
	c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	entityValue, err := u.service.Get(id)

	if err != nil {

		c.JSON(http.StatusNotFound, err.Error())

		return

	}
	model := FromModel(entityValue)

	c.JSON(http.StatusOK, model)
}
func (u *AbstractHandler[T, F]) GetAll(FromModel func(*T) *F, c *gin.Context) {

	allEntitys, err := u.service.GetAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	frontEndValues := []F{}

	for _, value := range allEntitys {
		newValue := FromModel(&value)
		frontEndValues = append(frontEndValues, *newValue)
	}
	c.JSON(http.StatusOK, frontEndValues)
}

func (u *AbstractHandler[T, F]) Create(model AbstactModel[T, F], FromModel func(*T) *F, c *gin.Context) {

	if err := c.BindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	entity := model.ToModel()

	err := u.service.Create(&entity)

	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, FromModel(&entity))
}

func (u *AbstractHandler[T, F]) Update(model AbstactModel[T, F], FromModel func(*T) *F, c *gin.Context) {

	if err := c.BindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	entity := model.ToModel()

	err := u.service.Update(&entity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, FromModel(&entity))
}

func (u *AbstractHandler[T, F]) Delete(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := u.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Usuario eliminado satisfactoriamente")

}
