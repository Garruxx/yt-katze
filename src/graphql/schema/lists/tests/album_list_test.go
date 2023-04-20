package tests

import (
	"katze/src/graphql/schema/common"
	"katze/src/graphql/schema/lists"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestAlbumList(t *testing.T) {

	// Test case 1, albums field
	albumsField := lists.AlbumType.Fields()["albums"]
	if albumsField == nil {
		t.Fatalf("Test case 1 failed, expected albums field no be no-nil")
	}
	// Test case 1.1 albums field type
	if albumsField.Type != common.AlbumsType {
		t.Fatalf(
			"Test case 1.1 failed, extected albums field type common.AlbumsType",
		)
	}

	// Test case 2, continuation field
	continuationField := lists.AlbumType.Fields()["continuation"]
	if continuationField == nil {
		t.Fatalf("Test case 2 failed, expected continuation field no be no-nil")
	}
	// Test case 2.1 continuation field type
	if continuationField.Type != common.ContinuationType {
		t.Fatalf(
			"Test case 2.1 failed, extected continuation field type common.ContinuationType",
		)
	}

	// Test case 3, continuationId field
	continuationIdField := lists.AlbumType.Fields()["continuationId"]
	if continuationIdField == nil {
		t.Fatalf(
			"Test case 3 failed, expected continuationId field no be no-nil",
		)
	}
	// Test case 3.1 continuationId field type
	if continuationIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 3.1 failed, extected continuationId field type graphql.String",
		)
	}

	// Test case 4, visitorId field
	visitorIdField := lists.AlbumType.Fields()["visitorId"]
	if visitorIdField == nil {
		t.Fatalf("Test case 4 failed, expected visitorId field no be no-nil")
	}
	// Test case 4.1 visitorId field type
	if visitorIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 4.1 failed, extected visitorId field type graphql.String",
		)
	}
}
