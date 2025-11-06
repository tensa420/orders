package order_paid

import (
	"order/internal/entity"
	"order/pkg/kafka_structure/order_paid"
)

func EntityOrderPaidToProto(o *entity.OrderPaid) order_paid.OrderPaid {
	return order_paid.OrderPaid{
		OrderUUID:       o.OrderUUID,
		EventUUID:       o.EventUUID,
		UserUUID:        o.UserUUID,
		TransactionUUID: o.TransactionUUID,
		PaymentMethod:   PtrToString(o.PaymentMethod),
	}
}
func PtrToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
