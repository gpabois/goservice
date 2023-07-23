package events

import (
	"errors"

	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQP_Subscription struct {
	deliveries <-chan amqp.Delivery
}

func (sub AMQP_Subscription) Pump(s Subscriber) result.Result[bool] {
	d := <-sub.deliveries
	res := s(d.Body)

	if res.IsSuccess() {
		d.Ack(false)
	} else {
		d.Reject((true))
	}

	return res
}

type eventExchange struct {
	routingKey string
	exchange   string
}

// AMQP Based Event Service
type AMQP struct {
	channel          *amqp.Channel
	exchangeBindings map[string]eventExchange // Bindings eventName -> exchangeName, routingKey
	queueBindings    map[string]string        // Bindings eventName -> queueName
}

func NewAMQP(channel *amqp.Channel) *AMQP {
	return &AMQP{channel: channel, exchangeBindings: make(map[string]eventExchange), queueBindings: make(map[string]string)}
}

// Define the exchange which will route the message
func (a *AMQP) ConfigurePublish(eventName string, exchangeName string, routingKey string) {
	a.exchangeBindings[eventName] = eventExchange{
		exchange:   exchangeName,
		routingKey: routingKey,
	}
}

// Define the queue which receives the events
func (a *AMQP) ConfigureSubscribe(eventName string, queueName string) {
	a.queueBindings[eventName] = queueName
}

func (a *AMQP) Publish(eventName string, payload any) result.Result[bool] {
	ex, ok := a.exchangeBindings[eventName]
	if !ok {
		return result.Result[bool]{}.Failed(errors.New("the event publishing is not configured"))
	}

	res := serde.Serialize(payload, "application/json")
	if res.HasFailed() {
		return result.Result[bool]{}.Failed(res.UnwrapError())
	}

	data := res.Expect()
	err := a.channel.Publish(ex.exchange, ex.routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})

	if err != nil {
		result.Result[bool]{}.Failed(err)
	}

	return result.Success(true)
}

func (a *AMQP) Subscribe(eventName string, subscriberName string) result.Result[Subscribtion] {
	queueName, ok := a.queueBindings[eventName]

	if !ok {
		return result.Result[Subscribtion]{}.Failed(errors.New("the event subscribing is not configured"))
	}

	deliveries, err := a.channel.Consume(queueName, subscriberName, false, false, false, false, nil)
	if err != nil {
		return result.Result[Subscribtion]{}.Failed(err)
	}

	return result.Success[Subscribtion](AMQP_Subscription{deliveries: deliveries})
}
