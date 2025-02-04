package order

import (
	"github.com/shopspring/decimal"
	"time"
)

type Status string

const (
	PENDING                        Status = "PENDING"
	BLOCK_BALANCE                  Status = "BLOCK_BALANCE"
	SEND_ORDER_TO_ANOTHER_EXCHANGE Status = "SEND_ORDER_TO_ANOTHER_EXCHANGE"
	WAIT_ORDER_RESULT              Status = "WAIT_ORDER_RESULT"
	DONE_SUCCESSFULLY              Status = "DONE_SUCCESSFULLY"
	FAILED                         Status = "FAILED"
	UNRESOLVED                     Status = "UNRESOLVED"
)

type OrderModel struct {
	ID          uint64
	UserId      string
	AskCurrency string
	AskAmount   decimal.Decimal
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
