package common

import "github.com/graphql-go/graphql"

var ContinuationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Continuation",
	Fields: graphql.Fields{
		"query": &graphql.Field{
			Type: graphql.String,
		},
		"params": &graphql.Field{
			Type: graphql.String,
		},
	},
})
