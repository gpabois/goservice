package integration_testing

import (
	"log"

	"github.com/gpabois/gostd/result"
	"github.com/ory/dockertest/v3"
)

type ResourcesArgs struct {
}

type Resources struct {
	pool      *dockertest.Pool
	resources []*dockertest.Resource
}

func (mngr *Resources) Cleanup() result.Result[bool] {
	for _, res := range mngr.resources {
		err := mngr.pool.Purge(res)
		if err != nil {
			return result.Failed[bool](err)
		}
	}

	return result.Success(true)
}

func WithResources[T any](args ResourcesArgs, f func(mngr *Resources) result.Result[T]) result.Result[T] {
	res := newResourcesManager(args)
	if res.HasFailed() {
		return result.Result[T]{}.Failed(res.UnwrapError())
	}

	mngr := res.Expect()
	defer mngr.Cleanup()
	return f(mngr)
}

func newResourcesManager(args ResourcesArgs) result.Result[*Resources] {
	pool, err := dockertest.NewPool("")
	log.Println("Starting a new docker pool")
	if err != nil {
		log.Panicln(err)
		return result.Result[*Resources]{}.Failed(err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Panicln(err)
		return result.Result[*Resources]{}.Failed(err)
	}

	log.Println("Successfully connected to the docker pool")
	return result.Success(&Resources{pool: pool})
}
