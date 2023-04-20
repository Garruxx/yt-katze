package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema"

	"github.com/graphql-go/graphql"
)

// ArtistsList is a graphql query for the artists list.
var ArtistsList = &graphql.Field{
	Type: schema.ArtistsList,
	Name: "artistsList",
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
	Resolve: resolvers.ArtistList,
}
