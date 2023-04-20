package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var ArtistAlbums = &graphql.Field{
	Type: common.CardItemsType,
	Name: "artistAlbums",
	Args: graphql.FieldConfigArgument{
		"artistId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"params": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"visitorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.ArtistAlbums,
}
