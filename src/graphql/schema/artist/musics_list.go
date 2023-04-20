package artist

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var MusicsListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArtistMusicsList",
	Fields: graphql.Fields{
		"songs": &graphql.Field{
			Type: common.SongsType,
		},
		"params": &graphql.Field{
			Type: graphql.String,
		},
		"continuationId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
