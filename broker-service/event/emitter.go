package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Emitter struct {
	connection *amqp.Connection
	exchange   string
}

func (emitter *Emitter) setup() error {
	channel, err := emitter.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	return declareExchange(channel)
}

func (emitter *Emitter) Push(event string, severity string) error {
	channel, err := emitter.connection.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	log.Println("Pushing to channel")

	err = channel.Publish(
		emitter.exchange, // exchange
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	log.Println("Sent!")

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection, exchangeName string) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
		exchange:   exchangeName,
	}

	err := emitter.setup()

	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
