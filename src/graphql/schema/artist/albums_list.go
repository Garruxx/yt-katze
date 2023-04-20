package artist

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var AlbumsListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArtistAlbumsList",
	Fields: graphql.Fields{
		"albums": &graphql.Field{
			Type: common.CardItemsType,
		},
		"params": &graphql.Field{
			Type: graphql.String,
		},
		"continuationId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
