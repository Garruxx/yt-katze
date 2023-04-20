package common

import "github.com/graphql-go/graphql"

var BestMatchType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BestMatch",
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"albumType": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnails": &graphql.Field{
			Type: ThumbnailsType,
		},
		"watchId": &graphql.Field{
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
	},
})
