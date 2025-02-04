package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mostafanoorpur/aban-task/internal/httputil"
	"github.com/mostafanoorpur/aban-task/internal/order"
	"github.com/shopspring/decimal"
	"net/http"
)

type createNewOrderRequest struct {
	UserId            string `json:"user_id" validate:"required"` // we need to authenticate user but for now we skip authentication
	AskCurrency       string `json:"ask_currency" validate:"required"`
	AskCurrencyAmount string `json:"ask_currency_amount" validate:"required"`
}

func (c *createNewOrderRequest) toDTO() *order.AddOrderDto {
	askAmount, err := decimal.NewFromString(c.AskCurrencyAmount)
	if err != nil {
		return nil
	}
	return &order.AddOrderDto{
		UserId:            c.UserId,
		AskCurrency:       c.AskCurrency,
		AskCurrencyAmount: askAmount,
	}
}

func (controller *OrderHttpController) createNewOrderHandler(ctx echo.Context) error {
	var request createNewOrderRequest

	err := ctx.Bind(&request)
	if err != nil {
		return err
	}

	if err = ctx.Validate(&request); err != nil {
		return err
	}

	data := request.toDTO()
	if data == nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("send valid data please"))
	}

	err = controller.orderSvc.NewOrder(ctx.Request().Context(), request.toDTO())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, httputil.NewMessageResponse("order created successfully"))
}
