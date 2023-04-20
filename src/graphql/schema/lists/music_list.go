package lists

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var MusicType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MusicsList",
	Fields: graphql.Fields{
		"songs": &graphql.Field{
			Type: common.SongsType,
		},
		"continuation": &graphql.Field{
			Type: common.ContinuationType,
		},
		"continuationId": &graphql.Field{
			Type: graphql.String,
		},
		"visitorId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
