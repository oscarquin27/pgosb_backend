package antares_handlers

import (
	api_models "fdms/cmd/api/models"
	"fdms/src/models"
	"fdms/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AntaresController struct {
	antaresService services.AntaresService
}

func NewAntaresController(antaresService services.AntaresService) *AntaresController {
	return &AntaresController{
		antaresService: antaresService,
	}
}

func (u *AntaresController) GetAll(c *gin.Context) {

	//time.Sleep(12000 * time.Millisecond)

	allAntares, err := u.antaresService.GetAll()

	if err != nil {
		if err == models.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	antaresJsonList := []api_models.AntaresJson{}

	for _, ant := range allAntares {
		newAntares := api_models.ModelToAntaresJson(ant)
		antaresJsonList = append(antaresJsonList, *newAntares)
	}
	c.JSON(http.StatusOK, antaresJsonList)
}
