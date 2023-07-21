package graphql

import (
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	graphql_flow "github.com/gpabois/goservice/graphql/flow"
	"github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
	"github.com/graphql-go/graphql"
)

type SchemaInstantiator struct {
	Query option.Option[QueryInstantiator]
}

type Field interface {
	FieldInstantiator | graphql.Field
}

type Object interface {
	ObjectInstantiator | graphql.ObjectConfig
}

type ObjectInstantiator = func(flow.Flow) graphql.ObjectConfig

func WithFields(options ...utils.Configurator[graphql.Fields]) utils.Configurator[graphql.ObjectConfig] {
	return func(obj *graphql.ObjectConfig) {
		fs := graphql.Fields{}
		utils.Configure(&fs, options)
		obj.Fields = fs
	}
}

func WithField(name string, field FieldInstantiator) utils.Configurator[graphql.Fields] {
	return func(fields *graphql.Fields) {
		field := graphql.Field{}
		utils.Configure(&field, options)
		(*fields)[name] = &field
	}
}

func DynObject(name string, options ...utils.Configurator[graphql.ObjectConfig]) ObjectInstantiator {
	return func(parent flow.Flow) graphql.ObjectConfig {

		objCfg := graphql.ObjectConfig{
			Name:   name,
			Fields: fs,
		}

		utils.Configure(&objCfg, options)

		return objCfg
	}
}

type FieldInstantiator = func(flow.Flow) graphql.Field

func WithArg(name string, cfg graphql.ArgumentConfig) func(field *graphql.Field) {
	return func(field *graphql.Field) {
		field.Args[name] = &cfg
	}
}

func DynField(typ graphql.Output, mods []chain.Module, options ...utils.Configurator[graphql.Field]) FieldInstantiator {
	// Field
	return func(parent flow.Flow) graphql.Field {
		resolver := func(params graphql.ResolveParams) (any, error) {
			flo := parent.Fork()
			ch := chain.Chain{}
			ch.Install(mods...)
			flo = graphql_flow.SetResolveParams(flo, params)
			return ch.Call(flo).IntoAnyTuple()
		}

		field := graphql.Field{
			Type:    typ,
			Resolve: resolver,
		}

		for _, opt := range options {
			opt(&field)
		}

		return field
	}
}
