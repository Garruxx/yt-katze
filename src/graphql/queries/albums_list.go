package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema"

	"github.com/graphql-go/graphql"
)

var AlbumsList = &graphql.Field{
	Type: schema.AlbumsList,
	Name: "albumsList",
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
	Resolve: resolvers.AlbumsList,
}
