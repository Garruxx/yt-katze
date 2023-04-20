package graphql

import (
	"katze/logger"
	"katze/src/graphql/queries"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"general":                queries.General,
		"album":                  queries.Album,
		"albumsList":             queries.AlbumsList,
		"albumsPagination":       queries.AlbumsPagination,
		"artistAlbums":           queries.ArtistAlbums,
		"artistMusicsList":       queries.ArtistMusicsList,
		"artistMusicsPagination": queries.ArtistMusicsPagination,
		"artistProfile":          queries.ArtistProfile,
		"artistSingles":          queries.ArtistSingles,
		"artistsList":            queries.ArtistsList,
		"artistsPagination":      queries.ArtistsPagination,
		"musicsList":             queries.MusicsList,
		"musicsPagination":       queries.MusicsPagination,
		"single":                 queries.Single,
	},
	Description: "Query for the GraphQL API",
})

func GraphqlHandler() (*handler.Handler, error) {

	var Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryType,
	})
	if err != nil {
		logger.Errorf("Error creating schema: %v", err)
	}

	// Crate GraphQL controler
	graphqlHandler := handler.New(&handler.Config{
		Schema:     &Schema,
		Playground: true,
		Pretty:     true,
	})

	return graphqlHandler, nil
}
