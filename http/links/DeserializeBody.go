package http_links

import (
	"io/ioutil"
	"reflect"

	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http/flow"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

type DeserializeBodyArgs struct {
	FallbackContentType option.Option[string]
}

type ReflectDeserializeBodyArgs struct {
	Type reflect.Type
	DeserializeBodyArgs
}

// Deserialize the http body based on the Content-Type header, into the endpoint request
// Order 102
func Reflect_DeserializedBody(args ReflectDeserializeBodyArgs) chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		httpRequest := http_flow.GetHttpRequest(flo)
		contentType := httpRequest.Header.Get("Content-Type")

		if contentType == "" {
			contentType = args.FallbackContentType.Expect()
		}

		anyReq := endpoint_flow.GetEndpointRequest[any](flo).Expect()

		req := reflect.New(reflect.TypeOf(anyReq))
		req.Elem().Set(reflect.ValueOf(anyReq))

		raw, err := ioutil.ReadAll(httpRequest.Body)
		if err != nil {
			return chain.Result{}.Failed(err)
		}

		res := serde.Reflect_DeserializeInto(raw, contentType, req.Interface())

		if res.HasFailed() {
			return chain.Result{}.Failed(res.UnwrapError())
		}

		flo = endpoint_flow.SetEndpointRequest(flo, req.Elem().Interface())

		return next(flo)
	}, 102)
}

// Deserialize the http body based on the Content-Type header, into the endpoint request
// Order 102
func DeserializeBody[T any](args DeserializeBodyArgs) chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		httpRequest := http_flow.GetHttpRequest(flo)
		contentType := httpRequest.Header.Get("Content-Type")

		if contentType == "" {
			contentType = args.FallbackContentType.Expect()
		}

		endpointRequest := endpoint_flow.GetEndpointRequest[T](flo).Expect()
		decodedRes := serde.DeserializeFromReaderInto(httpRequest.Body, contentType, &endpointRequest)

		if decodedRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(decodedRes.UnwrapError())
		}

		flo = endpoint_flow.SetEndpointRequest(flo, endpointRequest)

		return next(flo)
	}, 102)
}
