package postgres

import (
	"github.com/mostafanoorpur/aban-task/internal/order"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	ID                uint64 `gorm:"primaryKey"`
	UserId            string
	AskCurrency       string
	AskCurrencyAmount decimal.Decimal
	Status            string
	createdAt         time.Time `gorm:"<-:create"`
	UpdatedAt         time.Time
	DeletedAt         time.Time `gorm:"index"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) ConvertEntityToModel(model *order.OrderModel) {
	o.UserId = model.UserId
	o.AskCurrency = model.AskCurrency
	o.AskCurrencyAmount = model.AskAmount
	o.Status = string(model.Status)
}
