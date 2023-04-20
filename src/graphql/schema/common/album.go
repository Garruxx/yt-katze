package common

import "github.com/graphql-go/graphql"

var AlbumObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Album",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"artists": &graphql.Field{
			Type: ArtistsType,
		},
		"thumbnails": &graphql.Field{
			Type: ThumbnailsType,
		},
		"single": &graphql.Field{
			Type: graphql.Boolean,
		},
		"ep": &graphql.Field{
			Type: graphql.Boolean,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"explicit": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var AlbumsType = graphql.NewList(graphql.NewNonNull(AlbumObjectType))