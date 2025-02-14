package abstract_handler

import (
	logger "fdms/src/infrastructure/log"
	"fdms/src/utils/results"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AbstractCRUDService[T any] interface {
	Get(id int64, isTemplate bool) *results.ResultWithValue[*T]
	GetAll(isTemplate bool, params ...string) ([]T, *results.GeneralError)
	Create(value *T, isTemplate bool) *results.ResultWithValue[*T]
	Update(value *T, isTemplate bool) *results.ResultWithValue[*T]
	Delete(id int64, isTemplate bool) *results.Result
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

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	r := u.service.Get(id, isTemplate)

	if !r.IsSuccessful {

		logger.Warn().Err(r.Err.AssociateException()).
			Msgf("El get con Id:%d no fue exitoso", id)

		if r.Err.Code() == results.NotFoundErr {

			c.JSON(http.StatusNotFound, r.Err.Message())
			return

		}

		c.JSON(http.StatusInternalServerError, r.Err.Message())
		return

	}
	model := FromModel(r.Value)

	c.JSON(http.StatusOK, model)
}

func (u *AbstractHandler[T, F]) GetAll(FromModel func(*T) *F, c *gin.Context) {

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	allEntitys, err := u.service.GetAll(isTemplate)

	if err != nil {

		logger.Warn().Err(err.AssociateException()).
			Msg("Problemas ejecutando GetAll")

		if err.Code() == results.NotFoundErr {
			c.JSON(http.StatusOK, allEntitys)
		}

		c.JSON(http.StatusInternalServerError, err.Message())
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

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	if err := c.BindJSON(&model); err != nil {

		logger.Warn().Err(err).Msg("Error parseando la data")

		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	entity := model.ToModel()

	r := u.service.Create(&entity, isTemplate)

	if !r.IsSuccessful {

		logger.Warn().Err(r.Err.AssociateException()).
			Msg("Problemas ejecutando Create")

		c.JSON(http.StatusInternalServerError, r.Err.Message())
		return
	}

	returnValue := FromModel(r.Value)

	c.JSON(http.StatusOK, returnValue)
}

func (u *AbstractHandler[T, F]) Update(model AbstactModel[T, F], FromModel func(*T) *F, c *gin.Context) {

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	if err := c.BindJSON(&model); err != nil {

		logger.Warn().Err(err).Msg("error parseando datos")
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	entity := model.ToModel()

	r := u.service.Update(&entity, isTemplate)

	if !r.IsSuccessful {

		logger.Warn().Err(r.Err.AssociateException()).
			Msg("Problemas ejecutando Update")

		if r.Err.Code() == results.NotFoundErr {

			c.JSON(http.StatusNotFound, r.Err.Message())
			return
		}

		c.JSON(http.StatusInternalServerError, r.Err.Message())
		return
	}

	c.JSON(http.StatusOK, FromModel(r.Value))
}

func (u *AbstractHandler[T, F]) Delete(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	isTemplate := strings.Contains(c.Request.URL.Path, "template")

	r := u.service.Delete(id, isTemplate)

	if !r.IsSuccessful {

		logger.Warn().Err(r.Err.AssociateException()).
			Msg("Problemas ejecutando Update")

		if r.Err.Code() == results.NotFoundErr {

			c.JSON(http.StatusNotFound, r.Err.Message())
			return
		}

		c.JSON(http.StatusInternalServerError, r.Err.Message())
		return
	}
	c.JSON(http.StatusOK, "Usuario eliminado satisfactoriamente")

}
