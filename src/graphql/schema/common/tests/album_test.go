package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestAlbum(t *testing.T) {

	// Test case 1, title field
	titleField := common.AlbumObjectType.Fields()["title"]
	if titleField == nil {
		t.Fatalf("Test case 1 failed: expected title field to be non-nil")
	}
	// Test case 1.1, title field type
	if titleField.Type != graphql.String {
		t.Fatalf("Test case 1.1 failed: expected title field type to be string")
	}

	// Test case 2, id field
	idField := common.AlbumObjectType.Fields()["id"]
	if idField == nil {
		t.Fatalf("Test case 2 failed: expected id field to be non-nil")
	}
	// Test case 2.1, id field type
	if idField.Type != graphql.String {
		t.Fatalf("Test case 2.1 failed: expected id field type to be string")
	}

	// Test case 3, artists field
	artistsField := common.AlbumObjectType.Fields()["artists"]
	if artistsField == nil {
		t.Fatalf("Test case 3 failed: expected artists field to be non-nil")
	}
	// Test case 3.1, artists field type
	if artistsField.Type != common.ArtistsType {
		t.Fatalf(
			"Test case 3.1 failed: expected artists field type to be [Artist]",
		)
	}

	// Test case 4, thumbnails field
	thumbnailsField := common.AlbumObjectType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 4 failed: expected thumbnails field to be non-nil")
	}
	// Test case 4.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf("Test case 4.1 failed: expected thumbnails field type to be ThumbnailsType")
	}

	// Test case 5, single field
	singleField := common.AlbumObjectType.Fields()["single"]
	if singleField == nil {
		t.Fatalf("Test case 5 failed: expected single field to be non-nil")
	}
	// Test case 5.1, single field type
	if singleField.Type != graphql.Boolean {
		t.Fatalf("Test case 5.1 failed: expected single field type to be boolean")
	}

	// Test case 6, ep field
	epField := common.AlbumObjectType.Fields()["ep"]
	if epField == nil {
		t.Fatalf("Test case 6 failed: expected ep field to be non-nil")
	}
	// Test case 6.1, ep field type
	if epField.Type != graphql.Boolean {
		t.Fatalf("Test case 6.1 failed: expected ep field type to be boolean")
	}

	// Test case 7, year field
	yearField := common.AlbumObjectType.Fields()["year"]
	if yearField == nil {
		t.Fatalf("Test case 7 failed: expected year field to be non-nil")
	}
	// Test case 7.1, year field type
	if yearField.Type != graphql.Int {
		t.Fatalf("Test case 7.1 failed: expected year field type to be int")
	}

	// Test case 8, explicit field
	explicitField := common.AlbumObjectType.Fields()["explicit"]
	if explicitField == nil {
		t.Fatalf("Test case 8 failed: expected explicit field to be non-nil")
	}
	// Test case 8.1, explicit field type
	if explicitField.Type != graphql.Boolean {
		t.Fatalf("Test case 8.1 failed: expected explicit field type to be boolean")
	}
}
