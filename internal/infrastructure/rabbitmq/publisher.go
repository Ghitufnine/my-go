package rabbitmq

import (
	"context"

	"github.com/ghitufnine/my-go/internal/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	rabbit *Rabbit
}

func NewPublisher(
	rabbit *Rabbit,
) repository.EventPublisher {
	return &Publisher{
		rabbit: rabbit,
	}
}

func (p *Publisher) Publish(
	ctx context.Context,
	topic string,
	payload []byte,
) error {

	return p.rabbit.Channel.PublishWithContext(
		ctx,
		p.rabbit.Exchange,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)
}
