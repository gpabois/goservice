package middlewares_tests

import (
	"bytes"
	"net/http"

	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

type endpointRequest struct {
	Value bool `serde:"value"`
}

func NewHttpRequestFixtureWithBody[T any](body T) *http.Request {
	// Create a mocked request
	buf := bytes.NewBuffer(serde.Serialize(body, "application/json").Expect())
	r := result.From(http.NewRequest("GET", "goservice.local", buf)).Expect()
	r.Header.Set("Content-Type", "application/json")
	return &r
}
