package graphql_links

import (
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	graphql_flow "github.com/gpabois/goservice/graphql/flow"
	"github.com/gpabois/gostd/result"
	"github.com/graphql-go/graphql"

	gql "github.com/gpabois/goservice/graphql"
)

type SchemaInstatiator func(flo flow.Flow) result.Result[graphql.Schema]

// Call the schema
// Order 400
func CallSchema(schemaInstantiator SchemaInstatiator) chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		schemaRes := schemaInstantiator(flo)
		if schemaRes.HasFailed() {
			return chain.Result{}.Failed(schemaRes.UnwrapError())
		}
		schema := schemaRes.Expect()

		// Get the result
		requestRes := graphql_flow.GetRequest(flo).IntoResult(gql.NewMissingRequest())
		if requestRes.HasFailed() {
			return chain.Result{}.Failed(requestRes.UnwrapError())
		}
		request := requestRes.Expect()

		params := graphql.Params{Schema: schema, RequestString: request}
		r := graphql.Do(params)

		if r.HasErrors() {
			return chain.Result{}.Failed(gql.NewExecutionErrors(r.Errors...))
		}

		flo = graphql_flow.SetResponse(flo, r.Data)
		return next(flo)
	}, 400)
}
