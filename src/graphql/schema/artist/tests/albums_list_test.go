package tests

import (
	"katze/src/graphql/schema/artist"
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestAlbumsList(t *testing.T) {

	// Test case 1, albums field
	albumsField := artist.AlbumsListType.Fields()["albums"]
	if albumsField == nil {
		t.Fatalf("Test case 1 failed, expected albums field no be no-nil")
	}
	// Test case 1.1 albums field type
	if albumsField.Type != common.CardItemsType {
		t.Fatalf(
			"Test case 1.1 failed, extected albums field type AlbumsType",
		)
	}

	// Test case 2, continuationId field
	continuationIdField := artist.AlbumsListType.Fields()["continuationId"]
	if continuationIdField == nil {
		t.Fatalf("Test case 2 failed, expected continuationId field no be no-nil")
	}
	// Test case 2.1 continuationId field type
	if continuationIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 2.1 failed, extected continuationId field type graphql.String",
		)
	}

	// Test case 3, artistId field
	artistIdField := artist.AlbumsListType.Fields()["artistId"]
	if artistIdField == nil {
		t.Fatalf("Test case 3 failed, expected artistId field no be no-nil")
	}
	// Test case 3.1 artistId field type
	if artistIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 3.1 failed, extected artistId field type graphql.String",
		)
	}
}
