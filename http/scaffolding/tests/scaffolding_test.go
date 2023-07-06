package http_scaffolding_tests

import (
	"testing"

	"github.com/gpabois/goservice/chain"
	http_scaffolding "github.com/gpabois/goservice/http/scaffolding"
)

func Test_HttpModuleScaffolding(t *testing.T) {
	chain.NewChain().Install(http_scaffolding.ScaffoldHttpModule(
		// Enable body deserialization
		http_scaffolding.EnableBodyDeserialization(
			http_scaffolding.WithContentTypeFallback("application/json"),
		),
		// Enable route parameters injection into the endpoint request
		http_scaffolding.EnableRouteParamInjection(),
		// Enable authentication
		http_scaffolding.EnableAuthentication(
			http_scaffolding.WitHeaderBasedAuthentication(
				http_scaffolding.WithAuthenticationHeaderName("Authorization"),
			),
		),
	))
}
