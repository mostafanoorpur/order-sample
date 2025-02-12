package order_test

import (
	"context"
	"errors"
	"github.com/mostafanoorpur/order-sample/internal/order"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	dbRepo = new(order.MockRepository)
)

func TestNewOrderService(t *testing.T) {
	assert.NotNil(t, order.NewOrderService(dbRepo), "NewOrderService() should not return nil")
}

func TestOrderService_NewOrder(t *testing.T) {
	svc := order.NewOrderService(dbRepo)

	type fields struct {
		dbRepo order.Repository
	}
	type args struct {
		ctx   context.Context
		order *order.AddOrderDto
	}
	testCases := []struct {
		Name          string
		MockBehaviour func(dbRepo *order.MockRepository)
		Fields        fields
		Args          args
		OutputOrder   *order.OrderModel
		ExpectedError error
	}{
		{
			Name: "create_order_without_any_error",
			MockBehaviour: func(dbRepo *order.MockRepository) {
				expectedOrder := &order.OrderModel{
					UserId:      "1",
					AskCurrency: "ABAN",
					AskAmount:   decimal.RequireFromString("3"),
					Status:      "PENDING",
				}
				dbRepo.On("Save", context.Background(), expectedOrder).Return(nil).Once()
			},
			Fields: fields{
				dbRepo: dbRepo,
			},
			Args: args{
				ctx: context.Background(),
				order: &order.AddOrderDto{
					UserId:            "1",
					AskCurrency:       "ABAN",
					AskCurrencyAmount: decimal.RequireFromString("3"),
				},
			},
			OutputOrder: &order.OrderModel{
				UserId:      "1",
				AskCurrency: "ABAN",
				AskAmount:   decimal.RequireFromString("0.01"),
				Status:      "PENDING",
			},
			ExpectedError: nil,
		},
		{
			Name: "failed_to_create_order_on_save_db_error",
			MockBehaviour: func(dbRepo *order.MockRepository) {
				expectedOrder := &order.OrderModel{
					UserId:      "1",
					AskCurrency: "ABAN",
					Status:      "PENDING",
				}

				err := errors.New("db_error")

				dbRepo.On("Save", context.Background(), expectedOrder).Return(err).Once()
			},
			Fields: fields{
				dbRepo: dbRepo,
			},
			Args: args{
				ctx: context.Background(),
				order: &order.AddOrderDto{
					UserId:      "1",
					AskCurrency: "ABAN",
				},
			},
			OutputOrder:   &order.OrderModel{},
			ExpectedError: errors.New("db_error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.MockBehaviour(dbRepo)

			err := svc.NewOrder(tc.Args.ctx, tc.Args.order)
			assert.Equal(t, tc.ExpectedError, err)
			dbRepo.AssertExpectations(t)
		})
	}
}
