package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestArtist(t *testing.T) {

	// Test case 1, id field
	idField := common.ArtistObjectType.Fields()["id"]
	if idField == nil {
		t.Fatalf("Test case 1 failed: expected id field to be non-nil")
	}
	// Test case 1.1, id field type
	if idField.Type != graphql.String {
		t.Fatalf("Test case 1.1 failed: expected id field type to be string")
	}

	// Test case 2, name field
	nameField := common.ArtistObjectType.Fields()["name"]
	if nameField == nil {
		t.Fatalf("Test case 2 failed: expected name field to be non-nil")
	}
	// Test case 2.1, name field type
	if nameField.Type != graphql.String {
		t.Fatalf("Test case 2.1 failed: expected name field type to be string")
	}

	// Test case 3, thumbnails field
	thumbnailsField := common.ArtistObjectType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 3 failed: expected thumbnails field to be non-nil")
	}
	// Test case 3.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf("Test case 3.1 failed: expected thumbnails field type to be ThumbnailsType")
	}
}
