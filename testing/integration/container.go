package integration_testing

import (
	"time"

	"github.com/goombaio/namegenerator"
)

func newContainerName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}
