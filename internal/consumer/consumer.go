package consumer

import (
	"context"

	"github.com/IBM/sarama"
)

type ShipAssembledConsumerService interface {
	RunConsumer(ctx context.Context, topic []string, handler ShipAssembledConsumer) error
}
type ShipAssembledConsumer interface {
	Setup(sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
	Cleanup(sarama.ConsumerGroupSession) error
}
