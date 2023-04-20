package common

import "github.com/graphql-go/graphql"

var ArtistObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Artist",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnails": &graphql.Field{
			Type: ThumbnailsType,
		},
	},
})

var ArtistsType = graphql.NewList(graphql.NewNonNull(ArtistObjectType))
