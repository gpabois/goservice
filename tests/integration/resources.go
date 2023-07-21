package integration_tests

import (
	"log"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/gpabois/gostd/result"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rabbitmq/amqp091-go"
)

type ResourcesManagerArgs struct {
}

type ResourcesManager struct {
	pool      *dockertest.Pool
	resources []*dockertest.Resource
}

func (mngr *ResourcesManager) Cleanup() result.Result[bool] {
	for _, res := range mngr.resources {
		err := mngr.pool.Purge(res)
		if err != nil {
			return result.Failed[bool](err)
		}
	}

	return result.Success(true)
}

func NewResourcesManager(args ResourcesManagerArgs) result.Result[*ResourcesManager] {
	pool, err := dockertest.NewPool("")
	log.Println("Starting a new docker pool")
	if err != nil {
		return result.Result[*ResourcesManager]{}.Failed(err)
	}

	err = pool.Client.Ping()
	if err != nil {
		return result.Result[*ResourcesManager]{}.Failed(err)
	}

	log.Println("Successfully connected to the docker pool")
	return result.Success(&ResourcesManager{pool: pool})
}

type RabbitMQ_Args struct{}

const RabbitMQ_DefaultPort = 5672

func newContainerName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}

func WithRabbitMQ(mngr *ResourcesManager, args RabbitMQ_Args) result.Result[*amqp091.Connection] {
	opts := dockertest.RunOptions{
		Repository: "rabbitmq",
		Name:       newContainerName(),
		Tag:        "latest",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5672/tcp": {{HostIP: "", HostPort: "5672"}},
		},
	}

	resource, err := mngr.pool.RunWithOptions(&opts)

	log.Println("Starting a RabbitMQ container")
	if err != nil {
		return result.Result[*amqp091.Connection]{}.Failed(err)
	}

	mngr.resources = append(mngr.resources, resource)

	res := result.Result[*amqp091.Connection]{}

	if err := mngr.pool.Retry(func() error {
		conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
		res = result.Success(conn)
		return err
	}); err != nil {
		return result.Result[*amqp091.Connection]{}.Failed(err)
	}

	log.Println("Successfuly connected to RabbitMQ at 5672")

	return res
}
