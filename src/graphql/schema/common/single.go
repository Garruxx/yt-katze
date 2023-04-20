package common

import "github.com/graphql-go/graphql"

var SingleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Single",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"artists": &graphql.Field{
			Type: ArtistsType,
		},
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"thumbnails": &graphql.Field{
			Type: ThumbnailsType,
		},
		"explicit": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
