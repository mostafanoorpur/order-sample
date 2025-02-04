package order

import "context"

type OrderService struct {
	repo Repository
}

type Repository interface {
	Save(ctx context.Context, model *OrderModel) error
}

func NewOrderService(repo Repository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) NewOrder(ctx context.Context, orderDto *AddOrderDto) error {
	return o.repo.Save(ctx, orderDto.toNewOrderModel())
}

// order simple state machine
func (o *OrderService) WatchAndDoOrders(ctx context.Context, order *OrderModel) error {
	switch order.Status {
	case PENDING:
		// TODO : check if this order can do or not
		// 1 : check if user has balance for this order or not
		// 2 : block balance and set status of order in BLOCK_BALANCE mode
		return nil
	case BLOCK_BALANCE:
		// TODO : call balance domain to block balance
		return nil
	case SEND_ORDER_TO_ANOTHER_EXCHANGE:
		// TODO : call providers buy_from_exchange method to send order
		return nil
	case WAIT_ORDER_RESULT:
		// TODO : wait result
		return nil
	}

	return nil
}
