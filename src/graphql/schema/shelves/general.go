package shelves

import (
	"katze/src/graphql/schema/common"
	"katze/src/graphql/schema/lists"

	"github.com/graphql-go/graphql"
)

var GeneralType = graphql.NewObject(graphql.ObjectConfig{
	Name: "General",
	Fields: graphql.Fields{
		"bestMatch": &graphql.Field{
			Type: common.BestMatchType,
		},
		"tracks": &graphql.Field{
			Type: lists.MusicType,
		},
		"albums": &graphql.Field{
			Type: lists.AlbumType,
		},
		"artists": &graphql.Field{
			Type: lists.ArtistType,
		},
		"visitorId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
