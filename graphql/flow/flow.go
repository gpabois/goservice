package graphql_flow

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/graphql-go/graphql"
)

func SetRequest(in flow.Flow, request string) flow.Flow {
	in["Graphql.Request"] = request
	return in
}

func GetRequest(in flow.Flow) option.Option[string] {
	return flow.Lookup[string]("Graphql.Request", in)
}

func SetResponse(in flow.Flow, response any) flow.Flow {
	in["Graphql.Response"] = response
	return in
}

func GetResponse[T any](in flow.Flow) option.Option[T] {
	return flow.Lookup[T]("Graphql.Response", in)
}

func SetResolveParams(in flow.Flow, params graphql.ResolveParams) flow.Flow {
	in["Graphql.Resolve.Params"] = params
	return in
}

func GetResolveParams(in flow.Flow) option.Option[graphql.ResolveParams] {
	return flow.Lookup[graphql.ResolveParams]("Graphql.Resolve.Params", in)
}
