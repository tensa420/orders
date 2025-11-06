package producer

import (
	"context"
	"order/internal/entity"
)

type OrderPaidProducer interface {
	SendMessage(ctx context.Context, topic string, order *entity.OrderPaid) error
}
