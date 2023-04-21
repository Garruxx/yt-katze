package common

import "github.com/graphql-go/graphql"

var CardItemObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CardItem",
	Fields: graphql.Fields{
		"thumbnails": &graphql.Field{
			Type: ThumbnailsType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"params": &graphql.Field{
			Type: graphql.String,
		},
		"explicit": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var CardItemsType = graphql.NewList(graphql.NewNonNull(CardItemObjectType))
