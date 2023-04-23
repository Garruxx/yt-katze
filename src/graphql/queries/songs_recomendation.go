package queries

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var SongsRecomendation = &graphql.Field{
	Name: "songsRecomendation",
	Type: common.SongsType,
	Args: graphql.FieldConfigArgument{
		"songId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"visitorId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.SongsRecomendation,
}
