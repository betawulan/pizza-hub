package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/service"
)

type orderDelivery struct {
	orderService service.OrderService
}

func (o orderDelivery) order(e echo.Context) error {
	var order model.Order

	err := e.Bind(&order)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	res, err := o.orderService.Order(e.Request().Context(), order)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusAccepted, res)
}

func AddOrderRoute(orderService service.OrderService, e *echo.Echo) {
	handler := orderDelivery{
		orderService: orderService,
	}

	e.POST("/orders", handler.order)
}
