package lists

import (
	"katze/src/graphql/schema/common"

	"github.com/graphql-go/graphql"
)

var AlbumType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AlbumsList",
	Fields: graphql.Fields{
		"albums": &graphql.Field{
			Type: common.AlbumsType,
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
