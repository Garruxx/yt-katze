package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema"

	"github.com/graphql-go/graphql"
)

var ArtistProfile = &graphql.Field{
	Type: schema.ShelfArtist,
	Name: "artistProfile",
	Args: graphql.FieldConfigArgument{
		"artistId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"visitorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.ArtistProfile,
}
