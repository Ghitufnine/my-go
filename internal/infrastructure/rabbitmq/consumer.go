package rabbitmq

import (
	"context"
	"log"

	"github.com/ghitufnine/my-go/internal/infrastructure/mongo"
)

type Consumer struct {
	rabbit  *Rabbit
	logRepo *mongo.TransactionLogRepository
}

func NewConsumer(
	rabbit *Rabbit,
	logRepo *mongo.TransactionLogRepository,
) *Consumer {

	return &Consumer{
		rabbit:  rabbit,
		logRepo: logRepo,
	}
}

func (c *Consumer) Start(
	ctx context.Context,
	queue string,
	routingKey string,
) error {

	ch := c.rabbit.Channel

	_, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		queue,
		routingKey,
		c.rabbit.Exchange,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	go func() {

		for d := range msgs {

			err := c.logRepo.Insert(
				ctx,
				d.RoutingKey,
				string(d.Body),
			)

			if err != nil {

				log.Println("mongo insert failed", err)

				d.Nack(false, true)
				continue
			}

			d.Ack(false)
		}

	}()

	return nil
}
