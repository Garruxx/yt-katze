package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var ArtistSingles = &graphql.Field{
	Type: common.CardItemsType,
	Name: "artistSingles",
	Args: graphql.FieldConfigArgument{
		"artistId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"continuationId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"visitorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.ArtistSingles,
}
