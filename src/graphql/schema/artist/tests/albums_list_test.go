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

	// Test case 2, params field
	paramsField := artist.AlbumsListType.Fields()["params"]
	if paramsField == nil {
		t.Fatalf("Test case 2 failed, expected params field no be no-nil")
	}
	// Test case 2.1 params field type
	if paramsField.Type != graphql.String {
		t.Fatalf(
			"Test case 2.1 failed, extected params field type graphql.String",
		)
	}

	// Test case 3, continuation field
	continuationField := artist.AlbumsListType.Fields()["continuationId"]
	if continuationField == nil {
		t.Fatalf("Test case 3 failed, expected continuation field no be no-nil")
	}
	// Test case 3.1 continuation field type
	if continuationField.Type != graphql.String {
		t.Fatalf(
			"Test case 3.1 failed, extected continuation field type graphql.String",
		)
	}
}
