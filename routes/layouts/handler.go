package routes

import (
	layout_errors "fdms/domain/entities/layouts"
	layouts "fdms/domain/layouts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LayoutController struct {
	layoutService layouts.LayoutRepository
}

func NewLayoutController(layoutService layouts.LayoutRepository) *LayoutController {
	return &LayoutController{
		layoutService: layoutService,
	}
}

func (u *LayoutController) GetLayout(c *gin.Context) {

	entity := c.Param("entity")

	layout, err := u.layoutService.GetLayout(entity)

	if err != nil {
		if err == layout_errors.ErrorLayoutFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, layout)
	return
}
