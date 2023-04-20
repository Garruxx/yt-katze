package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema"

	"github.com/graphql-go/graphql"
)

var AlbumsPagination = &graphql.Field{
	Type: schema.AlbumsList,
	Name: "albumsPagination",
	Args: graphql.FieldConfigArgument{
		"continuationId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"visitorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.AlbumsPagination,
}
