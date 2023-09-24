package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/service"
)

type chefDelivery struct {
	chefService service.ChefService
}

func (c chefDelivery) addChef(e echo.Context) error {
	var chef model.Chef

	err := e.Bind(&chef)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	res, err := c.chefService.AddChef(e.Request().Context(), chef)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, res)
}

func AddChefRoute(chefService service.ChefService, e *echo.Echo) {
	handler := chefDelivery{
		chefService: chefService,
	}

	e.POST("/chef", handler.addChef)
}
