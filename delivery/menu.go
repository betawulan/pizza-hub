package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/pizza-hub/service"
)

type menuDelivery struct {
	menuService service.MenuService
}

func (m menuDelivery) getMenus(e echo.Context) error {

	menus, err := m.menuService.GetMenus(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, menus)
}

func AddMenuRoute(menuService service.MenuService, e *echo.Echo) {
	handler := menuDelivery{
		menuService: menuService,
	}

	e.GET("/menus", handler.getMenus)
}
