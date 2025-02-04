package order

import (
	"github.com/shopspring/decimal"
)

type AddOrderDto struct {
	UserId            string
	AskCurrency       string
	AskCurrencyAmount decimal.Decimal
	AskPrice          decimal.Decimal
}

func (a *AddOrderDto) toNewOrderModel() *OrderModel {
	return &OrderModel{
		UserId:      a.UserId,
		AskCurrency: a.AskCurrency,
		AskAmount:   a.AskCurrencyAmount,
		Status:      PENDING,
	}
}
