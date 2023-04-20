package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema"

	"github.com/graphql-go/graphql"
)

var MusicsList = &graphql.Field{
	Type: schema.MusicsList,
	Name: "musicsList",
	Args: graphql.FieldConfigArgument{
		"query": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"visitorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"params": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.MusicsList,
}
