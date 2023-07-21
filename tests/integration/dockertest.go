package integration_tests

import (
	"log"

	"github.com/gpabois/gostd/result"
	"github.com/ory/dockertest/v3"
	"github.com/rabbitmq/amqp091-go"
)

type DockerResourcesManager struct {
	pool      *dockertest.Pool
	resources []*dockertest.Resource
}

func (mngr *DockerResourcesManager) Cleanup() result.Result[bool] {
	for _, res := range mngr.resources {
		err := mngr.pool.Purge(res)
		if err != nil {
			return result.Failed[bool](err)
		}
	}

	return result.Success(true)
}

func NewDockerResourcesManager(name string) result.Result[*DockerResourcesManager] {
	pool, err := dockertest.NewPool(name)
	log.Println("Starting a new docker pool")
	if err != nil {
		return result.Result[*DockerResourcesManager]{}.Failed(err)
	}

	err = pool.Client.Ping()
	if err != nil {
		return result.Result[*DockerResourcesManager]{}.Failed(err)
	}

	log.Println("Successfully connected to the docker pool")
	return result.Success(&DockerResourcesManager{pool: pool})
}

type RabbitMQ_Args struct{}

const RabbitMQ_DefaultPort = 5672

func WithRabbitMQ(mngr *DockerResourcesManager, args RabbitMQ_Args) result.Result[*amqp091.Connection] {
	resource, err := mngr.pool.Run("rabbitmq", "latest", []string{})

	log.Println("Starting a RabbitMQ container")
	if err != nil {
		return result.Result[*amqp091.Connection]{}.Failed(err)
	}

	mngr.resources = append(mngr.resources, resource)

	res := result.Result[*amqp091.Connection]{}

	if err := mngr.pool.Retry(func() error {
		conn, err := amqp091.Dial("amqp://localhost:5672")
		log.Println("Successfuly connected to RabbitMQ at 5672")
		res = result.Success(conn)
		return err
	}); err != nil {
		return result.Result[*amqp091.Connection]{}.Failed(err)
	}

	return res
}
