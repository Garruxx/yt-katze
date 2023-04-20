package lists

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var ArtistType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArtistList",
	Fields: graphql.Fields{
		"artists": &graphql.Field{
			Type: common.ArtistsType,
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
