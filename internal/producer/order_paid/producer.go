package order_paid

import "github.com/IBM/sarama"

type OrderPaidProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewOrderPaidProducer(producer sarama.SyncProducer, topic string) *OrderPaidProducer {
	return &OrderPaidProducer{
		producer: producer,
		topic:    topic,
	}
}
