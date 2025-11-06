package consumer

import (
	"context"
	"errors"
	"log"
	"order/internal/consumer"
	"order/internal/service"
	"order/pkg/kafka_structure/ship_assembled"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

type ShipAssembledConsumer struct {
	consumer      sarama.ConsumerGroup
	consumerTopic []string
	service       service.OrderService
}

func NewShipAssembledConsumer(cons sarama.ConsumerGroup, topic []string) *ShipAssembledConsumer {
	return &ShipAssembledConsumer{
		consumer:      cons,
		consumerTopic: topic,
	}
}

func (c *ShipAssembledConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ShipAssembledConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ShipAssembledConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("kafka channel closed")
				return nil
			}

			shipAssembled := ship_assembled.ShipAssembled{}
			err := proto.Unmarshal(message.Value, &shipAssembled)
			if err != nil {
				log.Printf("failed to unmarshal shipAssembled message: %v", err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err = c.service.SetOrderStatusCompleted(ctx, shipAssembled.OrderUUID)
			if err != nil {
				log.Printf("Failed to complete order: %v", err)
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			log.Printf("kafka session was done")
			return nil
		}
	}
}
func (c *ShipAssembledConsumer) RunConsumer(ctx context.Context, topic []string, handler consumer.ShipAssembledConsumer) error {
	for {
		if err := c.consumer.Consume(ctx, topic, handler); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}
			log.Printf("Kafka consume error: %v", err)
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		log.Printf("Kafka consumer group rebalancing...")
	}
}
