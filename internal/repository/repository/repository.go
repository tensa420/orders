package repository

import (
	"context"

	"order/internal/entity"

	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order entity.Order) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO orders (order_uuid,user_uuid,part_uuids,total_price,transaction_uuid,payment_method,status) VALUES($1,$2,$3,$4,$5,$6,$7)", order.OrderUUID, order.UserUUID, order.PartsUUID, order.TotalPrice, order.TransactionUUID, order.PaymentMethod, order.Status)
	if err != nil {
		return errors.Wrap(err, "failed to create order")
	}
	return nil
}

func (r *OrderRepository) GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error) {
	res, err := r.pool.Query(ctx, "SELECT * FROM orders WHERE (order_uuid) = $1", orderUUID)
	if err != nil {
		return nil, errors.Wrap(err, "order not found")
	}
	var order entity.Order
	err = res.Scan(
		&order.OrderUUID,
		&order.UserUUID,
		&order.PartsUUID,
		&order.TotalPrice,
		&order.TransactionUUID,
		&order.PaymentMethod,
		&order.Status)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) PayOrder(ctx context.Context, info entity.PaymentInfo) error {
	_, err := r.pool.Exec(ctx, "UPDATE orders SET transaction_uuid = $1,payment_method = $2 WHERE order_uuid = $3 ", info.TransactionUUID, info.PaymentMethod, info.OrderUUID)
	if err != nil {
		return errors.Wrap(err, "failed to update order")
	}
	return nil
}

func (r *OrderRepository) CancelOrder(ctx context.Context, orderUUID string) error {
	_, err := r.pool.Exec(ctx, "UPDATE orders SET (status) = $1 WHERE order_uuid = $2", "CANCELLED", orderUUID)
	if err != nil {
		return errors.Wrap(err, "failed to cancel order")
	}
	return nil
}
