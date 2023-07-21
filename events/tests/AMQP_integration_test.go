//go:build integration
// +build integration

package events_tests

import (
	"testing"

	"github.com/gpabois/goservice/events"
	integration_tests "github.com/gpabois/goservice/tests/integration"
)

type TestEvent struct {
	Arg string
}

func Test_AMQP(t *testing.T) {
	mngr := integration_tests.NewResourcesManager(integration_tests.ResourcesManagerArgs{}).Expect()
	defer mngr.Cleanup()

	// Create a RabbitMQ instance
	amqpConn := integration_tests.WithRabbitMQ(mngr, integration_tests.RabbitMQ_Args{}).Expect()

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
}
