package shelves

import (
	"katze/src/graphql/schema/artist"
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var ArtistType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArtistShelf",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnails": &graphql.Field{
			Type: common.ThumbnailsType,
		},
		"musicsList": &graphql.Field{
			Type: artist.MusicsListType,
		},
		"albumsList": &graphql.Field{
			Type: artist.AlbumsListType,
		},
		"singlesList": &graphql.Field{
			Type: artist.SinglesListType,
		},
		"visitorId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
