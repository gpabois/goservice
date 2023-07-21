package graphql_links

import (
	"io/ioutil"

	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	graphql_flow "github.com/gpabois/goservice/graphql/flow"
	http_flow "github.com/gpabois/goservice/http/flow"
)

// Run the Graphql Request/Response with HTTP as a transport method
func WithHttpTransport() chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		r := http_flow.GetHttpRequest(flo)
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return chain.Result{}.Failed(err)
		}
		request := string(bytes)

		flo = graphql_flow.SetRequest(flo, request)
		return next(flo)
	}, 0)
}
