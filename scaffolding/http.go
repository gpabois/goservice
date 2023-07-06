package scaffolding

import (
	"net/http"

	"github.com/gpabois/goservice/chain"
	endpoint_modules "github.com/gpabois/goservice/endpoint/modules"
	http_transport "github.com/gpabois/goservice/http"
	http_modules "github.com/gpabois/goservice/http/modules"
)

func Http[Request any, Response any](endpointModule endpoint_modules.Endpoint[Request, Response], httpModule http_modules.Http, modules ...chain.Module) http.Handler {
	ch := chain.NewChain()
	ch = ch.Install(endpointModule, httpModule)
	ch = ch.Install(modules...)
	return http_transport.NewHandler(ch)
}
