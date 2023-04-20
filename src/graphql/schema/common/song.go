package common

import "github.com/graphql-go/graphql"

var SongObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
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
		"album": &graphql.Field{
			Type: graphql.String,
		},
		"albumId": &graphql.Field{
			Type: graphql.String,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
		"explicit": &graphql.Field{
			Type: graphql.Boolean,
		},
		"thumbnails": &graphql.Field{
			Type: ThumbnailsType,
		},
		"trackNumber": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var SongsType = graphql.NewList(graphql.NewNonNull(SongObjectType))
