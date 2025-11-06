package order_paid

import (
	"context"
	"log"
	"order/internal/entity"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

func (p *OrderPaidProducer) SendMessage(ctx context.Context, topic string, ord *entity.OrderPaid) error {
	protoMessage := EntityOrderPaidToProto(ord)
	marshalledMessage, err := proto.Marshal(&protoMessage)
	if err != nil {
		return err
	}

	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(ord.OrderUUID),
		Value: sarama.ByteEncoder(marshalledMessage),
	})
	log.Printf("Partition:%d,Offset:%d", partition, offset)
	return nil
}
