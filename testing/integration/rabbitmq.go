package integration_testing

import (
	"log"

	"github.com/gpabois/gostd/result"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ_Args struct{}

const RabbitMQ_DefaultPort = 5672

func SetupRabbitMQ(mngr *Resources, args RabbitMQ_Args) result.Result[*amqp091.Connection] {
	opts := dockertest.RunOptions{
		Repository:   "rabbitmq",
		Name:         newContainerName(),
		Tag:          "latest",
		ExposedPorts: []string{"5672"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5672/tcp": {{HostIP: "", HostPort: "5672"}},
		},
	}

	resource, err := mngr.pool.RunWithOptions(&opts)
	log.Println("Starting a RabbitMQ container")
	if err != nil {
		log.Panicln(err)
		return result.Result[*amqp091.Connection]{}.Failed(err)
	}

	mngr.resources = append(mngr.resources, resource)
	res := result.Result[*amqp091.Connection]{}
	if err := mngr.pool.Retry(func() error {
		conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
		res = result.Success(conn)
		return err
	}); err != nil {
		log.Panicln(err)
		return result.Result[*amqp091.Connection]{}.Failed(err)
	}

	log.Println("Successfuly connected to RabbitMQ at 5672")
	return res
}
