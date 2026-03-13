package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Exchange string
}

func New(
	url string,
	exchange string,
) (*Rabbit, error) {

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("rabbit dial: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("rabbit channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("exchange declare: %w", err)
	}

	return &Rabbit{
		Conn:     conn,
		Channel:  ch,
		Exchange: exchange,
	}, nil
}
