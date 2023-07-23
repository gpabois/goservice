//go:build integration
// +build integration

package events_tests

import (
	"testing"

	"github.com/gpabois/goservice/events"
	integration_testing "github.com/gpabois/goservice/testing/integration"
	"github.com/gpabois/gostd/result"
)

type TestEvent struct {
	Arg string
}

func Test_AMQP(t *testing.T) {
	integration_testing.WithResources(integration_testing.ResourcesArgs{}, func(resources *integration_testing.Resources) result.Result[bool] {
		// Create a RabbitMQ instance
		amqpConn := integration_testing.SetupRabbitMQ(resources, integration_testing.RabbitMQ_Args{}).Expect()

		// Configure the rabbitmq
		channel, err := amqpConn.Channel()
		if err != nil {
			panic(err)
		}
		defer channel.Close()

		// Configure exchange, to send our event
		if err := channel.ExchangeDeclare("exchange.test", "fanout", true, true, false, false, nil); err != nil {
			panic(err)
		}
		if err := channel.QueueDeclare("queue.test", false, false, true, false, nil); err != nil {
			panic(err)
		}
		if err := channel.QueueBind("queue.test", "", "exchange.test", false, nil); err != nil {
			panic(err)
		}

		ev := events.NewAMQP(channel)
		ev.ConfigurePublish("test", "exchange.test", "*")
		ev.ConfigureSubscribe("test", "queue.test")
		ev.Publish("test", TestEvent{Arg: "Hello world !"})
		sub := ev.Subscribe("test", "subscriber.test")

		return result.Success(true)
	})
}
