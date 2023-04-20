package shelves

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var AlbumType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AlbumShelf",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"artists": &graphql.Field{
			Type: common.ArtistsType,
		},
		"thumbnails": &graphql.Field{
			Type: common.ThumbnailsType,
		},
		"tracks": &graphql.Field{
			Type: common.SongsType,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"visitorId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
