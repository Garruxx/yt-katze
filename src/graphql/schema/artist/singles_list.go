package artist

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var SinglesListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArtistSinglesList",
	Fields: graphql.Fields{
		"singles": &graphql.Field{
			Type: common.CardItemsType,
		},
		"continuationId": &graphql.Field{
			Type: graphql.String,
		},
		"artistId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
