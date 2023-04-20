package common

import "github.com/graphql-go/graphql"

var ThumbnailObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Thumbnail",
	Fields: graphql.Fields{
		"width": &graphql.Field{
			Type: graphql.Int,
		},
		"height": &graphql.Field{
			Type: graphql.Int,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var ThumbnailsType = graphql.NewList(graphql.NewNonNull(ThumbnailObjectType))
