package layout_handlers

import (
	"fdms/src/models"
	"fdms/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LayoutController struct {
	layoutService services.LayoutService
}

func NewLayoutController(layoutService services.LayoutService) *LayoutController {
	return &LayoutController{
		layoutService: layoutService,
	}
}

func (u *LayoutController) GetLayout(c *gin.Context) {

	entity := c.Param("entity")

	layout, err := u.layoutService.Get(entity)

	if err != nil {
		if err == models.ErrorLayoutFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, layout)
}
