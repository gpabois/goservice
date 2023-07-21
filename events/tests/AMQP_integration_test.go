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
		// Configure exchange
		if err := channel.ExchangeDeclare("exchange.test", "fanout", true, true, false, false, nil); err != nil {
			panic(err)
		}
		defer channel.Close()

		ev := events.NewAMQP(channel)
		ev.ConfigurePublish("test", "exchange.test", "*")
		ev.Publish("test", TestEvent{Arg: "Hello world !"})

		return result.Success(true)
	})
}
