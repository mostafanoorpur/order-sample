package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mostafanoorpur/aban-task/internal/order"
)

type orderService interface {
	NewOrder(ctx context.Context, orderDto *order.AddOrderDto) error
}

func NewOrderHttpController(
	orderSvc orderService,
) *OrderHttpController {
	return &OrderHttpController{
		orderSvc: orderSvc,
	}
}

type OrderHttpController struct {
	orderSvc orderService
}

func (controller *OrderHttpController) RegisterRoutes(g *echo.Group) {
	postGroup := g.Group("/order")
	postGroup.POST("", controller.createNewOrderHandler)
}
