package order

import (
	"context"
	"log"
	"order/internal/client/converter"
	"order/internal/entity"
	"os"

	"github.com/google/uuid"
)

func (s *OrderService) PayOrder(ctx context.Context, orderUUID string, paymentMethod string) (string, error) {
	convertedPaymentMethod := converter.ConvertPaymentMethod(paymentMethod)

	ord, err := s.repo.GetOrder(ctx, orderUUID)
	if err != nil {
		return "", err
	}

	transactionUUID, err := s.paymClient.PayOrder(ctx, orderUUID, ord.UserUUID, paymentMethod)
	if err != nil {
		return "", err
	}

	eventUUID, _ := uuid.NewUUID()
	orderPaid := &entity.OrderPaid{
		TransactionUUID: transactionUUID,
		UserUUID:        ord.UserUUID,
		PaymentMethod:   ord.PaymentMethod,
		EventUUID:       converter.UUIDToString(eventUUID),
		OrderUUID:       ord.OrderUUID,
	}

	err = s.prod.SendMessage(ctx, os.Getenv("KAFKA_PRODUCER_TOPIC"), orderPaid)
	if err != nil {
		log.Printf("Failed to send order paid: %v", err)
	}

	err = s.repo.PayOrder(ctx, entity.PaymentInfo{
		TransactionUUID: transactionUUID,
		OrderUUID:       orderUUID,
		PaymentMethod:   convertedPaymentMethod,
	})
	if err != nil {
		return "", err
	}

	return transactionUUID, nil
}
