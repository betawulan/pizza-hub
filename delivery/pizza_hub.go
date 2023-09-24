package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/service"
)

type pizzaHubDelivery struct {
	pizzaHubService service.PizzaHubService
}

func (p pizzaHubDelivery) addChef(c echo.Context) error {
	var chef model.Chef

	err := c.Bind(&chef)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := p.pizzaHubService.AddChef(c.Request().Context(), chef)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (p pizzaHubDelivery) getMenus(c echo.Context) error {
	
	menus, err := p.pizzaHubService.GetMenus(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, menus)
}

func AddPizzaHubRoute(pizzaHubService service.PizzaHubService, e *echo.Echo) {
	handler := pizzaHubDelivery{
		pizzaHubService: pizzaHubService,
	}

	e.POST("/chef", handler.addChef)
	e.GET("/menus", handler.getMenus)
}
