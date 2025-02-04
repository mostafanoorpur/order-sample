package postgres

import (
	"context"
	"github.com/mostafanoorpur/aban-task/internal/order"
	"gorm.io/gorm"
)

func NewOrderPostgresRepository(db *gorm.DB) order.Repository {
	return &UserPostgresRepository{db: db}
}

type UserPostgresRepository struct {
	db *gorm.DB
}

func (u *UserPostgresRepository) Save(ctx context.Context, model *order.OrderModel) error {
	m := new(Order)
	m.ConvertEntityToModel(model)
	return u.db.WithContext(ctx).Create(m).Error
}
